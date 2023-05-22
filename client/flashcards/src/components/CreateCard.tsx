import {FormEvent, useState} from "react";
import {useNavigate} from "react-router-dom";
import {Card, Client} from "../Client";

export function CreateCard(props: {client: Client}) {
  let client: Client = props.client
  const navigate = useNavigate()

  async function handleSubmit(e: FormEvent<HTMLFormElement>) {
    e.preventDefault()
    let questionElem = document.getElementById("question") as HTMLInputElement
    let ansElem = document.getElementById("answer") as HTMLInputElement
    let boxElem = document.getElementById("box") as HTMLInputElement

    let boxNum = parseInt(boxElem.value)
    let card = await client.createCard(new Card(boxNum, questionElem.value, ansElem.value))
    navigate("/cards")
  }
  return <div className="d-flex justify-content-center" >
    <form onSubmit={(e) => handleSubmit(e)}>
      <div className="card">
        <div className="card-header">
          <h1>
            Create Card
          </h1>
        </div>
        <div className="m-1">
          <div className="form-group">
            <label htmlFor="question">Question</label>
            <input type="text" className="form-control" id="question" placeholder="Enter Question" />
          </div>
          <div className="form-group">
            <label htmlFor="answer">Answer</label>
            <input type="text" className="form-control" id="answer" placeholder="Enter Answer" />
          </div>
          <div className="form-group">
            <label htmlFor="box">WhichBox</label>
            <input type="text" className="form-control" id="box" placeholder="Enter Answer" />
          </div>
          <button type="submit" className="d-flex btn btn-success justify-content-center">Submit</button>
        </div>
      </div>
    </form>
  </div>
}

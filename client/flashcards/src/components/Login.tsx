import {FormEvent, useEffect, useState} from "react"
import {useNavigate} from "react-router-dom"
import {Card, Client, JwtToken} from "../Client"

import "../App.css"

const initialCards: Card[] = []
export default function Login(props: any) {
  let client: Client = props.client
  const [token, setToken] = useState(new JwtToken())
  const navigate = useNavigate()

  async function handleSubmit(e: FormEvent<HTMLFormElement>) {
    e.preventDefault()
    let emailElem = document.getElementById("email") as HTMLInputElement
    let passwdElem = document.getElementById("password") as HTMLInputElement
    let loginResponse = await client.login(emailElem.value, passwdElem.value)
    client.token = loginResponse
    if (client.token.Token != undefined) {
      setToken(loginResponse)
      props.setLoggedInFn(true)
      navigate("/cards")
      return
    }
    alert("wrong credentials")
  }

  return <div className="d-flex justify-content-center" >
    <form onSubmit={(e) => handleSubmit(e)}>
      <div className="card">
        <div className="card-header">
          <h1>
            Login
          </h1>
        </div>
        <div className="m-1">
          <div className="form-group">
            <label htmlFor="email">Email</label>
            <input type="email" className="form-control" id="email" placeholder="Enter Email" />
          </div>
          <div className="form-group">
            <label htmlFor="password">Answer</label>
            <input type="password" className="form-control" id="password" placeholder="Enter Password" />
          </div>
          <button type="submit" className="d-flex btn btn-success justify-content-center">Login</button>
        </div>
      </div>
    </form>
  </div>
}

    //<form onSubmit={(e) => handleSubmit(e)}>
    //  <div className="card">
    //    <div className="card-header">
    //      <h1>
    //        Create Card
    //      </h1>
    //    </div>
    //    <div className="m-1">
    //      <div className="form-group">
    //        <label htmlFor="question">Question</label>
    //        <input type="text" className="form-control" id="question" placeholder="Enter Question" />
    //      </div>
    //      <div className="form-group">
    //        <label htmlFor="answer">Answer</label>
    //        <input type="text" className="form-control" id="answer" placeholder="Enter Answer" />
    //      </div>
    //      <div className="form-group">
    //        <label htmlFor="box">WhichBox</label>
    //        <input type="text" className="form-control" id="box" placeholder="Enter Answer" />
    //      </div>
    //      <button type="submit" className="d-flex btn btn-success justify-content-center">Submit</button>
    //    </div>
    //  </div>
    //</form>

//    <button onClick={() => {
//      handlerListCards()
//    }}>
//    List Cards
//    </button>
//    {cards.map((card, id) => {
//      return (
//        <>
//        <div key={id}>
//        Question: {card.Question}
//        <br/>
//        Answer: {card.Answer}
//        </div>
//        </> 
//      )
//    })}
//

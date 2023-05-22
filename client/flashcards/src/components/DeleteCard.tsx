import {Card as card} from "../Client";
import {FormEvent, useEffect, useState} from "react";
import {useNavigate} from "react-router-dom";
import {Card, Client} from "../Client";

export function DeletableCard(props: {card: card, id: number, deleteCardCb: (id: number) => void}) {
  return (
    <div key={props.id}>
      <div className="card">
        <div className="card-body">
          <h5 className="card-title">Question</h5>
          <p className="card-text">{props.card.Question}</p>
          <h5 className="card-title">Answer</h5>
          <p className="card-text">{props.card.Answer}</p>

          <button onClick={() => {
            props.deleteCardCb(props.card.ID)
          }} className="btn btn-danger">Delete</button>
        </div>
      </div>
    </div>
  )
}

const initialCards: Card[] = []
export default function DeleteCard(props: {client: Client}) {
  let client = props.client as Client
  let navigate = useNavigate()
  const [cards, setCards] = useState(initialCards)

  function deleteCardCb(cardId: number) {
    client.deleteCard(cardId).then((card) => {
      alert("CARD WITH ID: " + card.ID + " WAS DELETED")
      navigate("/cards")
    })
  }

  useEffect(() => {
    if (client.token.Token == "") {
      navigate("/")
      return
    }
    client.listCards().then((list) => {
      setCards(list.Data)
    })
  }, [])

  function listCards() {
    if (cards != null) {
      return cards.map((card, id) => {
        return (
          <center key={id}>
            <DeletableCard card={card} id={id} deleteCardCb={deleteCardCb} ></DeletableCard>
          </center>
        )
      })
    }
  }

  return (
    <>
      {listCards()}
    </>
  )
}


import {useEffect, useState} from "react"
import {useNavigate} from "react-router-dom"
import {Card, Client} from "../Client"
import {Card as CardComponent} from "./Card"

const initialCards: Card[] = []
export default function Cards(props: {client: Client}) {
  let client = props.client as Client
  let navigate = useNavigate()
  const [cards, setCards] = useState(initialCards)

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
          <div className="d-flex justify-content-center" >
            <CardComponent card={card} id={id}></CardComponent>
          </div>
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


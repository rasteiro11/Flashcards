import {useState} from "react";
import {Card as card} from "../Client";
import "../App.css"

export function Card(props: {card: card, id: number}) {
  const [showAnswer, setShowAnswer] = useState(false)

  return (
    <div key={props.id}>
      <div className="card">
        <div className="flip-card">
          <div className="flip-card-inner">
            <div className="flip-card-front">
              <h1>Question</h1>
              <p>{props.card.Question}</p>
            </div>
            <div className="flip-card-back">
              <h1>Answer</h1>
              <p>{props.card.Answer}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}


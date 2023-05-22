export class ListCardsResponse {
  Pages: number
  Page: number
  Total: number
  Data: Card[]
  constructor() {
    this.Page = 0
    this.Pages = 0
    this.Data = []
    this.Total = 0
  }
}

export class Card {
  ID = 0
  CreatedAt = new Date
  UpdatedAt = new Date
  UserID = 0
  WhichBox: number
  Question: string
  Answer: string

  constructor(box: number, question: string, answer: string) {
    this.WhichBox = box
    this.Question = question
    this.Answer = answer
  }
}

export class JwtToken {
  Token: string
  Expiration: Date

  constructor() {
    this.Token = ""
    this.Expiration = new Date()
  }
}

var baseUrl = "http://localhost:6969/flashcard"
export class Client {
  token: JwtToken
  constructor() {
    this.token = new JwtToken()
  }

  async findOneCard(id: number): Promise<Card> {
    let idStr: string = id.toString()
    let req = await fetch(baseUrl + "/user/flashcard/" + idStr, {
      method: "GET",
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': "Bearer " + this.token.Token
      },
    })
    return req.json() as Promise<Card>
  }

  async deleteCard(id: number): Promise<Card> {
    let idStr: string = id.toString()
    let req = await fetch(baseUrl + "/user/flashcard/" + idStr, {
      method: "DELETE",
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': "Bearer " + this.token.Token
      },
    })
    return req.json() as Promise<Card>
  }

  async listCards(): Promise<ListCardsResponse> {
    let req = await fetch(baseUrl + "/user/flashcards", {
      method: "POST",
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': "Bearer " + this.token.Token
      },
      body: JSON.stringify({
        "perPage": 10,
        "page": 1
      })
    })
    return req.json() as Promise<ListCardsResponse>
  }

  async createCard(card: Card): Promise<Card> {
    let req = await fetch(baseUrl + "/user/flashcard", {
      method: "POST",
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': "Bearer " + this.token.Token
      },
      body: JSON.stringify(card)
    })
    return req.json() as Promise<Card>
    //call(baseUrl + "/user/flashcard", "POST", card, )
    //
  }

  async login(email: string, password: string): Promise<JwtToken> {
    let req = await fetch(baseUrl + "/auth/signin", {
      method: "POST",
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
      },
      body: JSON.stringify({
        "email": email,
        "password": password
      })
    })

    return req.json() as Promise<JwtToken>
  }
}

//async function call<B, R>(url: string, method: string, body?: T, header?: {[key: string]: string}): Promise<T> {
//  var res: Promise<R>
//  let reqConfig: RequestInit = {
//    method: method,
//    headers: header,
//    body: JSON.stringify(body)
//  }
//
//  if(header == null) {
//    reqConfig.headers = {
//      'Content-Type': 'application/json',
//      'Accept': 'application/json',
//    }
//  }
//
//  if(body != null) {
//    reqConfig.body = JSON.stringify(body)
//  }
// 
//  let req = await fetch(url, reqConfig)
//  return req.json() as Promise<T>
//}




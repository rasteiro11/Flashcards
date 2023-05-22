import {useState} from 'react';
import {Link, Route, Routes} from 'react-router-dom';
import {Client} from './Client';
import Cards from './components/Cards';
import {CreateCard} from './components/CreateCard';
import DeleteCard from './components/DeleteCard';
import Login from './components/Login';


function Header(props: {loggedIn: boolean}) {
  if (props.loggedIn) {
    return (
      <nav className="navbar navbar-expand-lg navbar-light bg-light">
        <a className="navbar-brand" href="#">Flashcards</a>
        <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
          <span className="navbar-toggler-icon"></span>
        </button>
        <div className="collapse navbar-collapse" id="navbarNav">
          <ul className="navbar-nav">
            <li className="nav-item"><Link className="nav-link" to='/cards'>Your Cards</Link></li>
            <li className="nav-item"><Link className="nav-link" to='/create'>Create Card</Link></li>
            <li className="nav-item"><Link className="nav-link" to='/delete'>Delete Card</Link></li>
          </ul>
        </div>
      </nav>
    )
  }
  return (
    <nav className="navbar navbar-expand-lg navbar-light bg-light">
      <a className="navbar-brand" href="#">Flashcards</a>
      <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span className="navbar-toggler-icon"></span>
      </button>
      <div className="collapse navbar-collapse" id="navbarNav">
        <ul className="navbar-nav">
          <li className='nav-item'><Link  className="nav-link" to='/'>Login</Link></li>
        </ul>
      </div>
    </nav>
  )
}


var client = new Client()
export default function App() {
  const [loggedIn, setLoggedIn] = useState(false)
  return (
    <>
      <div>
        <ul>
          <Header loggedIn={loggedIn} />
        </ul>
        <hr />
        <Routes>
          <Route path='/' element={<Login client={client} setLoggedInFn={setLoggedIn} />} />
          <Route path='/cards' element={<Cards client={client} />} />
          <Route path='/create' element={<CreateCard client={client} />} />
          <Route path='/delete' element={<DeleteCard client={client} />} />
        </Routes>
      </div>
    </>
  )
}



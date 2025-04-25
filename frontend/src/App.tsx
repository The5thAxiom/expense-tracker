import './App.css'
import './pages/About'
import ReactDOM from "react-dom/client"
import {BrowserRouter, Routes, Route} from "react-router-dom"
import Layout from './pages/Layout';
import About from './pages/About';
import Help from './pages/Help';
import Home from './pages/Home';
import Payments from './pages/Payments';


function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout/>}>
          <Route index element={<Home/>}/>
          <Route path="/help" element={<Help/>}/>
          <Route path="/about" element={<About/>}/>
          <Route path="/payments" element={<Payments/>}/>
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App

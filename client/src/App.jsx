import { BrowserRouter, Route, Routes } from "react-router-dom";
import "./App.css";
import Header from "./components/Header";
import Layout from "./components/Layout";
import Homepage from "./pages/Homepage";
import Login from "./pages/Login";

function App() {
  return (
    <>
<BrowserRouter>
        <Routes>
          <Route path="/" element={<Layout />}>
            <Route index element={<Homepage />}/>
            <Route path="/login" element={<Login />}/>

          </Route>
        </Routes>
      </BrowserRouter>


    </>
  );
}

export default App;

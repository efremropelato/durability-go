import { useState } from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import { Greet, GetESLfattoriale } from "../wailsjs/go/main/App";

function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [name, setName] = useState('');
    const updateName = (e: any) => setName(e.target.value);
    const updateResultText = (result: string) => setResultText(result);

    const [rsl, setRsl] = useState(100);
    const [fa, setFa] = useState(1);
    const [fb, setFb] = useState(1);
    const [fc, setFc] = useState(1);
    const [fd, setFd] = useState(1);
    const [fe, setFe] = useState(1);
    const [ff, setFf] = useState(1);
    const [fg, setFg] = useState(1);
    const [esl, setEsl] = useState(10);
    const updateEsl = (result: number) => setEsl(result);

    function greet() {
        Greet(name).then(updateResultText);
    }

    function getESL() {
        GetESLfattoriale(rsl,fa,fb,fc,fd,fe,ff,fb).then(updateEsl);
    }

    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo" />
            <div id="result" className="result">{resultText}</div>
            <div id="input" className="input-box">
                <input id="name" className="input" onChange={updateName} autoComplete="off" name="input" type="text" />
                <button className="btn" onClick={greet}>Greet</button>
            </div>
            {esl}
            <button className="btn" onClick={getESL}>GetESL</button>
        </div>
    )
}

export default App

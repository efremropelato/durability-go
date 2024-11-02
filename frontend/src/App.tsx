import { useState } from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import { GetESLSimple, GetESLEvolute } from "../wailsjs/go/main/App";

function App() {

    const [rsl, setRsl] = useState(100);
    const updateRsl = (e: any) => setRsl(parseFloat(e.target.value));
    const [fa, setFa] = useState(1);
    const updateFa = (e: any) => setFa(parseFloat(e.target.value));
    const [fb, setFb] = useState(1);
    const updateFb = (e: any) => setFb(parseFloat(e.target.value));
    const [fc, setFc] = useState(1);
    const updateFc = (e: any) => setFc(parseFloat(e.target.value));
    const [fd, setFd] = useState(1);
    const updateFd = (e: any) => setFd(parseFloat(e.target.value));
    const [fe, setFe] = useState(1);
    const updateFe = (e: any) => setFe(parseFloat(e.target.value));
    const [ff, setFf] = useState(1);
    const updateFf = (e: any) => setFf(parseFloat(e.target.value));
    const [fg, setFg] = useState(1);
    const updateFg = (e: any) => setFg(parseFloat(e.target.value));
    const [esl, setEsl] = useState(10);
    const updateEsl = (result: number) => setEsl(Math.round(result * 100) / 100);

    function getESL() {
        GetESLSimple(rsl,fa,fb,fc,fd,fe,ff,fg).then(updateEsl);
    }

    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo" />
            <div id="result" className="result">{rsl}*{fa}*{fb}*{fc}*{fd}*{fe}*{ff}*{fg} = {esl}</div>
            <div id="input" className="input-box">
                <input id="rsl" className="input" defaultValue={rsl} onChange={updateRsl} autoComplete="off" name="input" type="number" step="1"/><br/>
                <input id="fa" className="input" defaultValue={fa} onChange={updateFa} autoComplete="off" name="input" type="number" min="0.8" max="1.2" step="0.1"/><br/>
                <input id="fb" className="input" defaultValue={fb} onChange={updateFb} autoComplete="off" name="input" type="number" min="0.8" max="1.2" step="0.1"/><br/>
                <input id="fc" className="input" defaultValue={fc} onChange={updateFc} autoComplete="off" name="input" type="number" min="0.8" max="1.2" step="0.1"/><br/>
                <input id="fd" className="input" defaultValue={fd} onChange={updateFd} autoComplete="off" name="input" type="number" min="0.8" max="1.2" step="0.1"/><br/>
                <input id="fe" className="input" defaultValue={fe} onChange={updateFe} autoComplete="off" name="input" type="number" min="0.8" max="1.2" step="0.1"/><br/>
                <input id="ff" className="input" defaultValue={ff} onChange={updateFf} autoComplete="off" name="input" type="number" min="0.8" max="1.2" step="0.1"/><br/>
                <input id="fg" className="input" defaultValue={fg} onChange={updateFg} autoComplete="off" name="input" type="number" min="0.8" max="1.2" step="0.1"/><br/>
                <button className="btn" onClick={getESL}>GetESL</button>
            </div>
        </div>
    )
}

export default App

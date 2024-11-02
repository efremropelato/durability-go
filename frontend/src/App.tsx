import { useState, useEffect } from 'react'
import { GetESLSimple } from "../wailsjs/go/main/App";
// GetESLEvolute

// import reactLogo from './assets/react.svg'
// import viteLogo from '/vite.svg'

import './App.css'

import Layout from './Layout'
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Button } from "@/components/ui/button"


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
    GetESLSimple(rsl, fa, fb, fc, fd, fe, ff, fg).then(updateEsl);
  }

  useEffect(getESL, [rsl, fa, fb, fc, fd, fe, ff, fg])

  return (
    <Layout>
      <Card>
        <CardHeader>
          <CardTitle>Metodo fattoriale semplice</CardTitle>
          <CardDescription>Card Description</CardDescription>
        </CardHeader>
        <CardContent>
        <div id="input" className="input-box">
            <input id="rsl" className="input" defaultValue={rsl} onChange={updateRsl} autoComplete="off" name="rsl" type="number" step="1" /><br />
            <input id="fa" className="input" defaultValue={fa} onChange={updateFa} autoComplete="off" name="fa" type="number" min="0.8" max="1.2" step="0.1" /><br />
            <input id="fb" className="input" defaultValue={fb} onChange={updateFb} autoComplete="off" name="fb" type="number" min="0.8" max="1.2" step="0.1" /><br />
            <input id="fc" className="input" defaultValue={fc} onChange={updateFc} autoComplete="off" name="fc" type="number" min="0.8" max="1.2" step="0.1" /><br />
            <input id="fd" className="input" defaultValue={fd} onChange={updateFd} autoComplete="off" name="fd" type="number" min="0.8" max="1.2" step="0.1" /><br />
            <input id="fe" className="input" defaultValue={fe} onChange={updateFe} autoComplete="off" name="fe" type="number" min="0.8" max="1.2" step="0.1" /><br />
            <input id="ff" className="input" defaultValue={ff} onChange={updateFf} autoComplete="off" name="ff" type="number" min="0.8" max="1.2" step="0.1" /><br />
            <input id="fg" className="input" defaultValue={fg} onChange={updateFg} autoComplete="off" name="fg" type="number" min="0.8" max="1.2" step="0.1" /><br />
            <Button className="btn" onClick={getESL}>GetESL</Button>
          </div>
        </CardContent>
        <CardFooter>
        <div id="result" className="result">{rsl}*{fa}*{fb}*{fc}*{fd}*{fe}*{ff}*{fg} = {esl}</div>
        </CardFooter>
      </Card>
    </Layout>
  )
}

export default App

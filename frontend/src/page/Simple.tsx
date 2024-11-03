import { useState, useEffect } from 'react'
import Grid from '@mui/material/Grid2';
import Typography from '@mui/material/Typography';
import { TextField, Button } from '@mui/material';

import { GetESLSimple } from "../../wailsjs/go/main/App";
// GetESLEvolute

export default function MainGrid() {

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
        <>
            <Typography component="h2" variant="h6" sx={{ mb: 2 }}>
                Simple
            </Typography>
            <Grid
                container
                spacing={2}
                columns={12}
                sx={{ mb: (theme) => theme.spacing(2) }}
            >
                <Grid>
                    <TextField id="rsl" defaultValue={rsl} onChange={updateRsl} autoComplete="off" name="rsl" type="number" slotProps={{ htmlInput: { step: 1 } }} />
                    <TextField id="fa" defaultValue={fa} onChange={updateFa} autoComplete="off" name="fa" type="number" slotProps={{ htmlInput: { min: 0.8, max: 1.2, step: 0.1 } }} />
                    <TextField id="fb" defaultValue={fb} onChange={updateFb} autoComplete="off" name="fb" type="number" slotProps={{ htmlInput: { min: 0.8, max: 1.2, step: 0.1 } }} />
                    <TextField id="fc" defaultValue={fc} onChange={updateFc} autoComplete="off" name="fc" type="number" slotProps={{ htmlInput: { min: 0.8, max: 1.2, step: 0.1 } }} />
                    <TextField id="fd" defaultValue={fd} onChange={updateFd} autoComplete="off" name="fd" type="number" slotProps={{ htmlInput: { min: 0.8, max: 1.2, step: 0.1 } }} />
                    <TextField id="fe" defaultValue={fe} onChange={updateFe} autoComplete="off" name="fe" type="number" slotProps={{ htmlInput: { min: 0.8, max: 1.2, step: 0.1 } }} />
                    <TextField id="ff" defaultValue={ff} onChange={updateFf} autoComplete="off" name="ff" type="number" slotProps={{ htmlInput: { min: 0.8, max: 1.2, step: 0.1 } }} />
                    <TextField id="fg" defaultValue={fg} onChange={updateFg} autoComplete="off" name="fg" type="number" slotProps={{ htmlInput: { min: 0.8, max: 1.2, step: 0.1 } }} />
                    <Button onClick={getESL}>GetESL</Button>
                    <div id="result" className="result">{rsl}*{fa}*{fb}*{fc}*{fd}*{fe}*{ff}*{fg} = {esl}</div>
                </Grid>
            </Grid>
        </>
    );
}

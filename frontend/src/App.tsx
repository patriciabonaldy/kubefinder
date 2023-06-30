import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {Greet} from "../wailsjs/go/main/App";
import TableConfigMap from './components/Table';
import { Box } from '@mui/material';


interface Response {
    Name: string,
    NameSpace: string,
    Value: string
  }
  

function App() {
    const [resultText, setResultText] = useState<Response[]>([]);
    const [name, setName] = useState('');
    const updateName = (e: any) => setName(e.target.value);
    const updateResultText = (result: any) => {
        let obj: Response[] = JSON.parse(result);
        setResultText(obj);
    };

    function greet() {
        Greet(name).then(updateResultText);
    }

    console.log("resultText: ", resultText, typeof(resultText))

    // @ts-ignore
    return (
        <>
        <div id="App">
            <img src={logo} id="logo" alt="logo" width="50" height="60"/>
            <div id="input" className="input-box">
                <input id="name" className="input" onChange={updateName} autoComplete="off" name="input" type="text"/>
                <button className="btn" onClick={greet}>Greet</button>
            </div>
        </div>
        <br/>
        <Box>
            <TableConfigMap dataValue={resultText}/>
        </Box>
        </>
    )
}

export default App

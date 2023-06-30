import { useMemo, useState } from "react";
import "./App.css";
import { Greet } from "../wailsjs/go/main/App";
import logo from "./assets/images/logo-universal.png";

import { Box } from "@mui/material";
import MaterialReactTable, { MRT_ColumnDef } from "material-react-table";

interface Response {
  Name: string;
  NameSpace: string;
  Value: string;
}

function App() {
  const [resultText, setResultText] = useState<Response[]>([]);
  const [name, setName] = useState("");
  const updateName = (e: any) => setName(e.target.value);
  const updateResultText = (result: any) => {
    let obj: Response[] = JSON.parse(result);
    setResultText(obj);
  };

  function greet() {
    Greet(name).then(updateResultText);
  }

  console.log("resultText: ", resultText, typeof resultText);

  const columnsData: MRT_ColumnDef<{
    Name: string;
    NameSpace: string;
    Value: string;
  }>[] = useMemo(
    () => [
      {
        header: "Name",
        accessorKey: "Name",
      },
      {
        header: "NameSpace",
        accessorKey: "NameSpace",
      },
      {
        header: "Value",
        accessorKey: "Value",
      },
    ],
    []
  );

  // @ts-ignore
  return (
    <>
      <Box justifyContent="center" alignItems="center">
        <img src={logo} id="logo" alt="logo" />
        <br />
        <div id="input" className="input-box">
          <input
            id="name"
            className="input"
            onChange={updateName}
            autoComplete="off"
            name="input"
            type="text"
          />
          <button className="btn" onClick={greet}>
            Find
          </button>
        </div>
      </Box>

      <br />
      {resultText.length > 0 ? (
        <Box m="40px 0 0 0">
            <Box>
                <MaterialReactTable columns={columnsData} data={resultText} />
            </Box>
        </Box>
      ) : (
        <p>Not data exists</p>
      )}
    </>
  );
}

export default App;

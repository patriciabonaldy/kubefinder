import { useMemo } from "react";
import MaterialReactTable, { MRT_ColumnDef, MRT_Row } from "material-react-table";
import { Box } from "@mui/material";


export default function TableConfigMap(dataValue:any) {
  const columnsData: MRT_ColumnDef<{ Name: string; NameSpace: string; Value: string; }>[] = useMemo(
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
      }
    ],
    []
  );

  const dataMock = [
    {
      "Name": "fanscore-cricket-handler-prod",
      "NameSpace": "fanscore-core",
      "Value": "fanscore-poll-upsert-queue-prod",
    },
    {
      "Name": "fanscore-cricket-handler-stage",
      "NameSpace": "fanscore-core",
      "Value": "fanscore-poll-upsert-queue-stage",
    },
  ];
  console.log(dataValue)
  return (
    <Box>
      <MaterialReactTable columns={columnsData} data={dataValue} />
    </Box>
  );
}

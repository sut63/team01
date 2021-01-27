import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';

import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntPrescription } from '../../api/models/EntPrescription';

const useStyles = makeStyles(theme => ({
  table: {
    minWidth: 650,
  },
  formControl: {
    margin: theme.spacing(3),
    width: 200,
  },
}));



export default function ComponentsTable(sim: any) {
  const classes = useStyles();
  const api = new DefaultApi();
  const [Prescriptions, setPrescriptions] = useState<EntPrescription[]>(Array);


  const [loading, setLoading] = useState(true);

  const getPrescriptions = async () => {
    const res = await api.listPrescription({ limit: 100, offset: 0 });
    setLoading(false);
    setPrescriptions(res);
  };


  //console.log(Prescriptions)


  useEffect(() => {

    getPrescriptions();


  }, [loading]);



  var p = 0;
  console.log(sim.sim)
  for (var val of Prescriptions) {
    if (val.edges?.prescriptionpatient?.id === sim.sim || sim.sim === 0) {
      p = p + 1
    }
    //console.log(val)
  }
  console.log("p = ", p,"sim",sim.sim)
  if (p !== 0) {
    return (

      <TableContainer component={Paper}>
        <Table className={classes.table} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell align="center">No.</TableCell>
              <TableCell align="center">ชื่อ ผู้ป่วย</TableCell>
              <TableCell align="center">ชื่อแพทย์ที่สั่งยา</TableCell>
              <TableCell align="center">ชื่อยา</TableCell>
              <TableCell align="center">จำนวนยา</TableCell>
              <TableCell align="center">อาการ</TableCell>
              <TableCell align="center">หมายเหตุ</TableCell>
            </TableRow>
          </TableHead>
          
          <TableBody>
            {Prescriptions === undefined
              ? null
              : Prescriptions.filter(i => i.edges?.prescriptionpatient?.id === sim.sim || sim.sim === 0).map((item: any) => (
                <TableRow>
                  <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.edges?.prescriptionpatient?.name}</TableCell>
                  <TableCell align="center">{item.edges?.prescriptiondoctor?.name}</TableCell>
                  <TableCell align="center">{item.edges?.prescriptionmedicine?.name}</TableCell>
                  <TableCell align="center">{item.value}</TableCell>
                  <TableCell align="center">{item.symptom}</TableCell>
                  <TableCell align="center">{item.annotation}</TableCell>
                  
                </TableRow>
              ))}
          </TableBody>
        </Table>
      </TableContainer>
    );
  } else {
    return (
      <TableContainer component={Paper}>
        <Table className={classes.table} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell align="center">ไม่พบรายการการสั่งจ่ายยาของผู้ป่วยที่ท่านค้นหา กรุณาเพิ่มรายการสั่งจ่ายยา</TableCell>
            </TableRow>

          </TableHead>

        </Table>
      </TableContainer>
    );
  }



}
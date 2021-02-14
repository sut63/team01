import React, { FC, useEffect, useState } from 'react';
import Swal from 'sweetalert2'; // alert

import { Content, Header, Page, pageTheme } from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';

import {
  FormControl,
  TextField,
  Button,
  Typography,
  Grid,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Avatar,
} from '@material-ui/core';

import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import {
  EntDispenseMedicine,
  EntPrescription,
  EntBill,
} from '../../api/models';

import { Cookies } from 'react-cookie/cjs'; //cookie

// css style
const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    flexRow: {
      display: 'flex',
      flexWrap: 'wrap',
      flexDirection: 'row',
      justifyContent: 'center',
    },

    flexRowNoCen: {
      display: 'flex',
      flexWrap: 'wrap',
      flexDirection: 'row',
      justifyContent: 'space-between',
    },

    headLabel: {
      margin: theme.spacing(1),
    },

    formControl: {
      margin: 'auto',
      maxWidth: 350,
      minWidth: 260,
    },

    cardMargin: {
      marginBottom: theme.spacing(2),
    },

    tableMargin: {
      margin: theme.spacing(1),
    },

    button: {
      marginTop: theme.spacing(3),
      marginLeft: theme.spacing(1),
      marginRight: theme.spacing(1),
      width: 110,
      height: 50,
    },

    AlertMargin: {
      marginBottom: theme.spacing(5),
      width: 300,
    },
  }),
);

const Bill: FC<{}> = () => {
  const classes = useStyles();
  const profile = { givenName: 'ค้นหารายการการชำระค่ายา' };
  const cookies = new Cookies();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [statusTable, setStatusTable] = useState(false);
  const [table1, setable1] = useState(true);
  const [table2, setable2] = useState(false);

  const [billid, setbillid] = React.useState<number>(0);
  //Sweet Alert
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 5000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

  //structure receive data from api

  const [apibill, setApiBills] = useState<EntBill[]>([]);
  const [sapibill, ssetApiBills] = useState<EntBill>();

  const [apiprescriptions, setApiPrescription] = useState<EntPrescription[]>(
    [],
  );
  const [apidispensemedicine, setApiDispenseMedicine] = useState<
    EntDispenseMedicine[]
  >([]);

  //structure show data
  const [sPharmacistID, setPharmacistID] = useState(Number);
  const [sPharmacistName, setPharmacistName] = useState(String);

  //function get from api

  const getDispenseMedicine = async () => {
    const res = await api.listDispensemedicine({ limit: 0, offset: 0 });
    setLoading(false);
    setApiDispenseMedicine(res);
  };

  const getBill = async () => {
    const res = await api.listBill({ limit: 0, offset: 0 });
    setLoading(false);
    setApiBills(res);
  };

  const getPrescription = async () => {
    const res = await api.listPrescription({ limit: 8, offset: 0 });
    setLoading(false);
    setApiPrescription(res);
  };

  const checkPosition = async () => {
    if (cookies.get('PositionData') != 'Bill') {
      history.pushState('', '', '/' + cookies.get('PositionData'));
      window.location.reload(false);
    }
    setPharmacistID(Number(cookies.get('ID')));
    setPharmacistName(cookies.get('Name'));
    console.log(sPharmacistID);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getDispenseMedicine();
    getPrescription();
    checkPosition();
    getBill();
  }, [loading]);

  // handleChange
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,
  ) => {
    const name = event.target.name as keyof typeof Bill;
    const { value } = event.target;
    setbillid(value);
    console.log(billid);
  };

  

  //function SetData selectPrescriptions

  // clear input form

  const SearchBill = async () => {
    if (billid > 0) {
      
        const apiUrl = 'http://localhost:8080/api/v1/bills/' + billid;
        const requestOptions = {
          method: 'GET',
          headers: { 'Content-Type': 'application/json' },
        };
        
        fetch(apiUrl, requestOptions)
          .then(response => response.json())
          .then(async data => {
            console.log(data);
            ssetApiBills(data.data);
            console.log(sapibill);
            if (data.status === true) {
              Toast.fire({
                icon: 'success',
                title: 'ค้นหาข้อมูลสำเร็จ',
              });
              console.log(sapibill?.edges?.Payments?.pay);
            setable1(false);
            setable2(true);
            } else {
              Toast.fire({
                icon: 'error',
                title: 'ไม่พบข้อมูลประวัติการจ่ายยาของผู้ป่วยที่ค้นหา',
              });
              setable1(false);
              setable2(false);
    
            }
            
            
          });
        

      /* } else if (validateCardNumberError != '') {
      Toast.fire({
        icon: 'error',
        title: validateCardNumberError,
      });
      setStatusTable(false); */
    } else if (billid === 0) {
      Toast.fire({
        icon: 'info',
        title: 'ข้อมูลประวัติการชำระค่ายาทั้งหมด',
      });
      setStatusTable(false);
      setable1(true);
      setable2(false);
    } else {
      setStatusTable(false);
    }
  };

  return (
    <Page theme={pageTheme.website}>
      <Header title={`${profile.givenName}`}>
        <Avatar src="/broken-image.jpg" />
        <div style={{ marginLeft: 10, color: 'white' }}>
          <Typography variant="h6">{sPharmacistName}</Typography>
        </div>
      </Header>

      <Content>
        <Grid
          container
          direction="row"
          justify="space-evenly"
          alignItems="flex-start"
        >
          <Grid item xs={2}>
            <FormControl
              variant="outlined"
              className={classes.formControl}
              fullWidth
            >
              <Typography className={classes.headLabel}>
                เลขที่ใบจ่ายยา
              </Typography>
              <TextField
                label="กรุณาใส่เลขที่ใบจ่ายยา"
                name="amount"
                value={billid || ''}
                variant="outlined"
                type="number"
                onChange={handleChange}
              ></TextField>
            </FormControl>
          </Grid>

          <Grid item xs={2}>
            {/* <FormControl
                    variant="outlined"
                    className={classes.formControl}
                    fullWidth
                  >
                    <Typography className={classes.headLabel}>ชื่อคนไข้</Typography>
                    <TextField
                      label="กรุณาใส่ชื่อคนไข้"
                      name="amount"
                      value = {sBill.amount}
                      variant="outlined"
                      onChange={handleChange2}
                    ></TextField>
                    
                  </FormControl> */}
          </Grid>
          <Grid item xs={8}></Grid>
          <Grid item xs={1}>
            <Button
              variant="contained"
              color="secondary"
              className={classes.button}
              onClick={() => {
                SearchBill();
              }}
            >
              <Typography variant="button">ค้นหา</Typography>
            </Button>
          </Grid>
          <Grid item xs={11}></Grid>
          <Grid item xs>
            <Typography variant="h5" style={{ marginBottom: 25 }}>
              ตารางข้อมูล
            </Typography>

            <TableContainer component={Paper}>
              <Table aria-label="simple table">
                <TableHead>
                  <TableRow>
                    <TableCell align="center">No.</TableCell>
                    <TableCell align="center">เวลา</TableCell>
                    <TableCell align="center">รายชื่อคนไข้</TableCell>
                    <TableCell align="center">ช่องทางการชำระ</TableCell>
                    <TableCell align="center">จำนวนเงิน</TableCell>
                    <TableCell align="center">หมายเหตุ</TableCell>
                  </TableRow>
                </TableHead>
                {table1 ? (
                  <TableBody>
                    {apibill.map(item =>
                      apidispensemedicine
                        .filter(t => t.id === item.edges?.dispenseMedicines?.id)
                        .map(item3 =>
                          apiprescriptions
                            .filter(y => y.id === item3.edges?.prescription?.id)
                            .map(item2 => (
                              <TableRow key={item.id}>
                                <TableCell align="center">{item.id}</TableCell>
                                <TableCell align="center">
                                  {item3.datetime}
                                </TableCell>
                                <TableCell align="center">
                                  {item2.edges?.prescriptionpatient?.name}
                                </TableCell>
                                <TableCell align="center">
                                  {item.edges?.payments?.pay}
                                </TableCell>
                                <TableCell align="center">
                                  {item.amount}
                                </TableCell>
                                <TableCell align="center">
                                  {item.annotation}
                                </TableCell>
                              </TableRow>
                            )),
                        ),
                    )}
                  </TableBody>
                ) : null}
                {table2 ? (
                  <TableBody>
                    <TableCell align="center">{sapibill?.id}</TableCell>
                    <TableCell align="center">
                      {sapibill?.edges?.DispenseMedicines?.datetime}
                    </TableCell>
                    <TableCell align="center">
                      {
                        sapibill?.edges?.DispenseMedicines?.edges?.Prescription
                          ?.edges?.Prescriptionpatient?.name
                      }
                    </TableCell>
                    <TableCell align="center">
                      {sapibill?.edges?.Payments?.pay}
                    </TableCell>
                    <TableCell align="center">{sapibill?.amount}</TableCell>
                    <TableCell align="center">{sapibill?.annotation}</TableCell>
                  </TableBody>
                ) : null}
              </Table>
            </TableContainer>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
};
export default Bill;

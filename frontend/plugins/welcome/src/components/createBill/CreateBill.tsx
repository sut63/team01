import React, { FC, useEffect, useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import Swal from 'sweetalert2'; // alert

import { Content, Header, Page, pageTheme } from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';

import {
  FormControl,
  MenuItem,
  TextField,
  Button,
  Typography,
  Grid,
  Card,
  CardContent,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Avatar,
  Select,
} from '@material-ui/core';

import SaveIcon from '@material-ui/icons/Save';
import DeleteIcon from '@material-ui/icons/Delete';
import { Alert } from '@material-ui/lab'; //alert

import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import {
  EntDispenseMedicine,
  EntPrescription,
  EntPatientInfo,
  EntPayment,
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

//structure data change
interface Bill {
  amount: number;
  annotation: string;
  payer: string;
  dispensemedicine: number;
  payment: number;
}

const Bill: FC<{}> = () => {
  const classes = useStyles();
  const profile = { givenName: 'บันทึกการชำระค่ายา' };
  const cookies = new Cookies();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [amountError, setamountError] = React.useState('');
  const [payerError, setpayerError] = React.useState('');
  const [annotationError, setannotationError] = React.useState('');
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
  const [sBill, setBill] = React.useState<
    Partial<Bill>
  >({});

  const [apipayment, setPayment] = useState<EntPayment[]>([]);
  const [apibill, setApiBills] = useState<EntBill[]>([]);
  
  const [apiprescriptions, setApiPrescription] = useState<EntPrescription[]>(
    [],
  );
  const [apidispensemedicine, setApiDispenseMedicine] = useState<
    EntDispenseMedicine[]
  >([]);

  //structure show data
  const [sPharmacistID, setPharmacistID] = useState(Number);
  const [sPharmacistName, setPharmacistName] = useState(String);
  const [sNamePatient, setNamePatient] = useState(String);

  //function get from api
  const getPayment = async () => {
    const res = await api.listPayment({ limit: 0, offset: 0 });
    setLoading(false);
    setPayment(res);
  };

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
    if(cookies.get('PositionData') != 'Bill'){
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
    getPayment();
    getBill();
  }, [loading]);

  // handleChange
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof Bill;
    const { value } = event.target;
    const validateValue = value as string
    checkPattern(name, validateValue)
    setBill({ ...sBill, [name]: value });
    console.log(sBill);
  };

  const handleChange2 = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof Bill;
    const { value } = event.target;
    const validateValue = value as string
    checkPattern(name, validateValue)
    setBill({ ...sBill, [name]: Number(value) });
    console.log(sBill);
  };

  //map apiprescriptions for filter data have in apidispensemedicine
  const DispensemedicineMap = apidispensemedicine.filter(
    presc =>
      !apibill.some(
        dispe => dispe.edges?.dispenseMedicines?.id === presc.id,
      ),
  );
  //check and Alart massge
  const validatename = (val: string) => {
    return val != null ? true : false;
  }
  const validateannotation = (val: string) => {
    return val != null ? true : false;
  }
  const validateamount = (val: string) => {
    return val.charCodeAt(0) != 48 ? true : false ;
  }
  const checkPattern  = (id: string, value: string) => {
    switch(id) {
      case 'Payer':
        validatename(value) ? setpayerError('') : setpayerError('กรุณากรอกชื่อผู้จ่ายเงิน');
        return;
      case 'annotation':
        validateannotation(value) ? setannotationError('') : setannotationError('กรุณากรอกหมายเหตุ ถ้าไม่มีใส่ -');
        return;
      case 'amount':
        validateamount(value) ? setamountError('') : setamountError('กรุณากรอกค่ารักษา หรือราคาต้องไม่ติดลบ')
        return;
      default:
        return;
    }
  }

  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }
  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'payer':
        alertMessage("error","กรุณากรอกชื่อผู้จ่ายเงิน");
        return;
      case 'amount':
        alertMessage("error","กรุณากรอกค่ารักษา หรือราคาต้องไม่ติดลบ");
        return;
      case 'annotation':
        alertMessage("error","กรุณากรอกหมายเหตุ ถ้าไม่มีใส่ -");
        return;
      default:
        alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }

  //function SetData selectPrescriptions
  function selectPrescriptions(id: any, namePatient: any) {
    setBill({ ...sBill, dispensemedicine: Number(id) });
    setNamePatient(String(namePatient));
    console.log(sBill);
  }

  // clear input form
  function clear() {
    setBill({});
    setNamePatient('');
  }


  // function CreateDispenseMedicines data
  /* const CreateDispenseMedicines = async () => {
    if (
      sBill.annotation != undefined &&
      sPharmacistID != 0 &&
      sBill.payment != 0 &&
      sBill.dispensemedicine != 0 &&
      sBill.amount != 0
    ) {
      const Bills = {
        amount: sBill.amount,
        annotation: sBill.annotation,
        payer: sBill.payer,
        dispenseMedicine: sBill.dispensemedicine,
        payment: sBill.payment,
        pharmacist: sPharmacistID,
        };

      const res: any = await api.createBill({
        bill: Bills,
      });
      console.log(sBill);

      setStatus(true);
      if (res.id != '') {
        setAlert(true);
        setLoading(true);
        clear();
      } else {
        setAlert(false);
      }
    } else {
      setStatus(true);
      setAlert(false);
    }
    setTimeout(() => {
      setStatus(false);
    }, 5000);
  }; */

  const Create_Bill = async () => {
    const Bills = {
      amount: sBill.amount,
      annotation: sBill.annotation,
      payer: sBill.payer,
      dispenseMedicine: sBill.dispensemedicine,
      payment: sBill.payment,
      pharmacist: sPharmacistID,
      };
    const apiUrl = 'http://localhost:8080/api/v1/bills';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(Bills),
    };
  
    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(async data => {
        console.log(data);
        if (data.status === true) {
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          checkCaseSaveError(data.error.Name);

        }
      });
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
                    <TableCell align="center">เลือก</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {DispensemedicineMap.map(item => (apiprescriptions.filter(t => t.id === item.edges?.prescription?.id).map(item2 => (
                    <TableRow key={item.id}>
                      <TableCell align="center">{item.id}</TableCell>
                      <TableCell align="center">{item.datetime}</TableCell>
                      <TableCell align="center">{item2.edges?.prescriptionpatient?.name}</TableCell>
                      <TableCell align="center">
                        <Button
                          onClick={() => {
                            selectPrescriptions(
                              item.id,
                              item2.edges?.prescriptionpatient?.name,
                            );
                          }}
                          style={{ marginLeft: 2 }}
                          variant="contained"
                          color="primary"
                        >
                          เลือก
                        </Button>
                      </TableCell>

                    </TableRow>
                  ))))}
                </TableBody>
              </Table>
            </TableContainer>
          </Grid>

          <Grid
            item
            xs={6}
            style={{
              minWidth: 400,
            }}
          >
            <Grid className={classes.flexRowNoCen}>
              <Typography variant="h5" style={{ marginBottom: 25 }}>
                ข้อมูลใบจ่ายยา
              </Typography>
              {status ? (
                <div>
                  {alert ? (
                    <Alert
                      variant="filled"
                      severity="success"
                      style={{ fontSize: 16 }}
                    >
                      บันทึกข้อมูลสำเร็จ!
                    </Alert>
                  ) : (
                    <Alert
                      variant="filled"
                      severity="warning"
                      style={{ fontSize: 16 }}
                    >
                      บันทึกข้อมูลไม่สำเร็จ!
                    </Alert>
                  )}
                </div>
              ) : null}
            </Grid>
            <Card className={classes.cardMargin}>
              <CardContent>
                <Grid className={classes.flexRow}>
                  <FormControl
                    variant="outlined"
                    className={classes.formControl}
                    style={{
                      maxWidth: 200,
                      minWidth: 100,
                    }}
                  >
                    <Typography className={classes.headLabel}>
                      ID
                    </Typography>
                    <TextField
                      label="ID"
                      name="dispensemedicine"
                      variant="outlined"
                      onChange={handleChange}
                      value={sBill.dispensemedicine || ''} // (undefined || '') = ''
                      inputProps={{ readOnly: true }}
                    ></TextField>
                  </FormControl>

                  <FormControl
                    variant="outlined"
                    className={classes.formControl}
                    fullWidth
                  >
                    <Typography className={classes.headLabel}>คนไข้</Typography>
                    <TextField
                      label="ชื่อ-นามสกุล"
                      name="name"
                      variant="outlined"
                      onChange={handleChange}
                      value={sNamePatient || ''} // (undefined || '') = ''
                      inputProps={{ readOnly: true }}
                    ></TextField>
                  </FormControl>
                </Grid>
              </CardContent>
            </Card>

          

            <Card className={classes.cardMargin}>
              <CardContent>
                <Grid className={classes.flexRow}>
                <FormControl
                    variant="outlined"
                    className={classes.formControl}
                    fullWidth
                  >
                    <Typography className={classes.headLabel}>หมายเหตุ</Typography>
                    <TextField
                      label="หมายเหตุ"
                      name="annotation"
                      value = {sBill.annotation}
                      onChange={handleChange}
                      variant="outlined"
                    ></TextField>
                  </FormControl>
                  <FormControl
                    variant="outlined"
                    className={classes.formControl}
                    fullWidth
                  >
                    <Typography className={classes.headLabel}>ราคา</Typography>
                    <TextField
                      label="ราคา"
                      name="amount"
                      value = {sBill.amount}
                      variant="outlined"
                      onChange={handleChange2}
                    ></TextField>
                  </FormControl>

                  <FormControl
                    variant="outlined"
                    className={classes.formControl}
                    fullWidth
                  >
                    <Typography className={classes.headLabel}>
                    ช่องทางการชำระ
                    </Typography>

                    <Select
                      labelId="Payment"
                      label="ช่องทางการชำระ"
                      name="payment"
                      value={sBill.payment || ''} // (undefined || '') = ''
                      onChange={handleChange}
                    >
                      {apipayment.map(option => (
                        <MenuItem key={option.id} value={option.id}>
                          {option.pay}
                        </MenuItem>
                      ))}
                    </Select>
                  </FormControl>

                  <FormControl
                    variant="outlined"
                    className={classes.formControl}
                    fullWidth
                  >
                    <Typography className={classes.headLabel}>ผู้จ่ายเงิน</Typography>
                    <TextField
                      label="ผู้จ่ายเงิน"
                      name="payer"
                      value = {sBill.payer}
                      variant="outlined"
                      onChange={handleChange}
                    ></TextField>
                  </FormControl>

                  <FormControl
                    variant="outlined"
                    className={classes.flexRow}
                    fullWidth
                  >
                    <Button
                      variant="contained"
                      color="secondary"
                      className={classes.button}
                      startIcon={<SaveIcon />}
                      onClick={() => {
                        Create_Bill();
                      }}
                    >
                      <Typography variant="button">บันทึก</Typography>
                    </Button>

                    <Button
                      variant="contained"
                      className={classes.button}
                      startIcon={<DeleteIcon />}
                      onClick={() => {
                        clear();
                      }}
                    >
                      <Typography variant="button">ล้างค่า</Typography>
                    </Button>
                  </FormControl>
                </Grid>
              </CardContent>
            </Card>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
};
export default Bill;

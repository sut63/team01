import React, { FC, useEffect, useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';

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
  dispensemedicine: number;
  payment: number;
}

const Bill: FC<{}> = () => {
  const classes = useStyles();
  const profile = { givenName: 'บันทึกการจ่ายยา' };
  const cookies = new Cookies();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);

  //structure receive data from api
  const [sBill, setBill] = React.useState<
    Partial<Bill>
  >({});

  const [apipayment, setPayment] = useState<EntPayment[]>([]);
  
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

  const getPrescription = async () => {
    const res = await api.listPrescription({ limit: 8, offset: 0 });
    setLoading(false);
    setApiPrescription(res);
  };

  const checkPosition = async () => {
    setPharmacistID(Number(cookies.get('ID')));
    setPharmacistName(cookies.get('Name'));
  };

  // Lifecycle Hooks
  useEffect(() => {
    getDispenseMedicine();
    getPrescription();
    checkPosition();
    getPayment();
  }, [loading]);

  // handleChange
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof Bill;
    const { value } = event.target;
    setBill({ ...sBill, [name]: value });
    console.log(sBill);
  };

  //map apiprescriptions for filter data have in apidispensemedicine
  const prescriptionMap = apiprescriptions.filter(
    presc =>
      !apidispensemedicine.some(
        dispe => dispe.edges?.prescription?.id === presc.id,
      ),
  );

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
  const CreateDispenseMedicines = async () => {
    if (
      sBill.annotation != undefined &&
      sPharmacistID != 0 &&
      sBill.payment != 0 &&
      sBill.dispensemedicine != 0 &&
      sBill.amount != 0
    ) {
      const Bills = {
        annotation: sBill.annotation,
        pharmacist: sPharmacistID,
        payment: sBill.payment,
        dispensemedicine: sBill.dispensemedicine,
        amount: sBill.amount
      };

      const res: any = await api.createBill({
        bill: Bills,
      });

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
    }, 2000);
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
              ตารางข้อมูลคิว
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
                  {apidispensemedicine.map(item => (apiprescriptions.filter(t => t.id === item.edges?.prescription?.id).map(item2 => (
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
                ข้อมูลใบสั่งยา
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
                      onChange={handleChange}
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
                    className={classes.flexRow}
                    fullWidth
                  >
                    <Button
                      variant="contained"
                      color="secondary"
                      className={classes.button}
                      startIcon={<SaveIcon />}
                      onClick={() => {
                        CreateDispenseMedicines();
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

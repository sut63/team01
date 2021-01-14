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
  EntAnnotation,
  EntDispenseMedicine,
  EntPrescription,
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
interface DispenseMedicine {
  annotation: number;
  datetime: string;
  prescription: number;
}

const DispenseMedicine: FC<{}> = () => {
  const classes = useStyles();
  const profile = { givenName: 'บันทึกการจ่ายยา' };
  const cookies = new Cookies();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);

  //structure receive data from api
  const [sDispensemedicine, setDispenseMedicine] = React.useState<
    Partial<DispenseMedicine>
  >({});
  const [apiannotations, setApiAnnotation] = useState<EntAnnotation[]>([]);
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
  const [sNameDoctor, setNameDoctor] = useState(String);

  //function get from api
  const getAnnotation = async () => {
    const res = await api.listAnnotation({ limit: 0, offset: 0 });
    setLoading(false);
    setApiAnnotation(res);
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
    getAnnotation();
    getDispenseMedicine();
    getPrescription();
    checkPosition();
  }, [loading]);

  // handleChange
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof DispenseMedicine;
    const { value } = event.target;
    setDispenseMedicine({ ...sDispensemedicine, [name]: value });
  };

  //map apiprescriptions for filter data have in apidispensemedicine
  const prescriptionMap = apiprescriptions.filter(
    presc =>
      !apidispensemedicine.some(
        dispe => dispe.edges?.prescription?.id === presc.id,
      ),
  );

  //function SetData selectPrescriptions
  function selectPrescriptions(id: any, namePatient: any, nameDoctor: any) {
    setDispenseMedicine({ ...sDispensemedicine, prescription: Number(id) });
    setNamePatient(String(namePatient));
    setNameDoctor(String(nameDoctor));
  }

  // clear input form
  function clear() {
    setDispenseMedicine({});
    setNamePatient('');
    setNameDoctor('');
  }

  // function CreateDispenseMedicines data
  const CreateDispenseMedicines = async () => {
    if (
      sDispensemedicine.annotation != undefined &&
      sDispensemedicine.datetime != undefined &&
      sDispensemedicine.datetime != ':00+07:00' &&
      sPharmacistID != 0 &&
      sDispensemedicine.prescription != undefined
    ) {
      const dispensemedicines = {
        annotation: sDispensemedicine.annotation,
        datetime: sDispensemedicine.datetime + ':00+07:00',
        pharmacist: sPharmacistID,
        prescription: sDispensemedicine.prescription,
      };

      const res: any = await api.createDispensemedicine({
        dispensemedicine: dispensemedicines,
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
                    <TableCell align="center">รายชื่อคนไข้</TableCell>
                    <TableCell align="center">เลือกคิว</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {prescriptionMap.map((item: any) => (
                    <TableRow key={item.id}>
                      <TableCell align="center">{item.id}</TableCell>
                      <TableCell align="center">
                        {item.edges?.prescriptionpatient?.name}
                      </TableCell>
                      <TableCell align="center">
                        <Button
                          onClick={() => {
                            selectPrescriptions(
                              item.id,
                              item.edges?.prescriptionpatient?.name,
                              item.edges?.prescriptiondoctor?.name,
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
                  ))}
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
                      ลำดับที่
                    </Typography>
                    <TextField
                      label="เลขคิว"
                      name="queue"
                      variant="outlined"
                      value={sDispensemedicine.prescription || ''} // (undefined || '') = ''
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
                      value={sNamePatient || ''} // (undefined || '') = ''
                      inputProps={{ readOnly: true }}
                    ></TextField>
                  </FormControl>

                  <FormControl
                    variant="outlined"
                    className={classes.formControl}
                    fullWidth
                  >
                    <Typography className={classes.headLabel}>
                      แพทย์ผู้ตรวจ
                    </Typography>
                    <TextField
                      label="ชื่อ-นามสกุล"
                      name="physician"
                      variant="outlined"
                      value={sNameDoctor || ''} // (undefined || '') = ''
                      inputProps={{ readOnly: true }}
                    ></TextField>
                  </FormControl>
                </Grid>
              </CardContent>
            </Card>

            <Card className={classes.cardMargin}>
              <CardContent>
                <Grid className={classes.flexRow}>
                  <TableContainer
                    component={Paper}
                    className={classes.tableMargin}
                  >
                    <Table aria-label="simple table">
                      <TableHead>
                        <TableRow>
                          <TableCell align="center">No.</TableCell>
                          <TableCell align="center">ชื่อยา</TableCell>
                          <TableCell align="center">หมายเลข</TableCell>
                          <TableCell align="center">แบนด์</TableCell>
                          <TableCell align="center">จำนวน</TableCell>
                          <TableCell align="center">
                            จำนวนที่เหลือในคลัง
                          </TableCell>
                        </TableRow>
                      </TableHead>
                      <TableBody>
                        {apiprescriptions
                          .filter(
                            item => item.id === sDispensemedicine.prescription,
                          )
                          .map((item: any) => (
                            <TableRow key={item.id}>
                              <TableCell align="center">
                                {item.edges?.prescriptionmedicine?.id}
                              </TableCell>
                              <TableCell align="center">
                                {item.edges?.prescriptionmedicine?.name}
                              </TableCell>
                              <TableCell align="center">
                                {item.edges?.prescriptionmedicine?.serial}
                              </TableCell>
                              <TableCell align="center">
                                {item.edges?.prescriptionmedicine?.brand}
                              </TableCell>
                              <TableCell align="center">{item.value}</TableCell>
                              <TableCell align="center">
                                {item.edges?.prescriptionmedicine?.amount -
                                  item.value}
                              </TableCell>
                            </TableRow>
                          ))}
                      </TableBody>
                    </Table>
                  </TableContainer>
                </Grid>
              </CardContent>
            </Card>

            <Card className={classes.cardMargin}>
              <CardContent>
                <Grid className={classes.flexRow}>
                  <FormControl
                    variant="outlined"
                    className={classes.formControl}
                  >
                    <Typography className={classes.headLabel}>
                      วันที่
                    </Typography>
                    <TextField
                      label="วันที่"
                      name="datetime"
                      type="datetime-local"
                      value={sDispensemedicine.datetime || ''} // (undefined || '') = ''
                      onChange={handleChange}
                      InputLabelProps={{
                        shrink: true,
                      }}
                    />
                  </FormControl>

                  <FormControl
                    variant="outlined"
                    className={classes.formControl}
                    fullWidth
                  >
                    <Typography className={classes.headLabel}>
                      หมายเหตุ
                    </Typography>

                    <Select
                      labelId="annotation"
                      label="เลือกหมายเหตุ"
                      name="annotation"
                      value={sDispensemedicine.annotation || ''} // (undefined || '') = ''
                      onChange={handleChange}
                    >
                      {apiannotations.map(option => (
                        <MenuItem key={option.id} value={option.id}>
                          {option.messages}
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
export default DispenseMedicine;

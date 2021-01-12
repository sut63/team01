import React, { FC, useEffect, useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';

import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';

import {
  Container,
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
} from '@material-ui/core';

import SaveIcon from '@material-ui/icons/Save';
import { Alert } from '@material-ui/lab'; //alert

import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import {
  EntAnnotation,
  EntPrescription,
} from '../../api/models';

// css style
const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    flexRow: {
      display: 'flex',
      flexWrap: 'wrap',
      flexDirection: 'row',
      justifyContent: 'center',
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
    },
  }),
);

const DispenseMedicine: FC<{}> = () => {
  const classes = useStyles();
  const profile = { givenName: 'บันทึกการจ่ายยา' };

  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);

  //structure receive data from api
  const [annotations, setAnnotation] = useState<EntAnnotation[]>([]);
  const [prescriptions, setPrescription] = useState<EntPrescription[]>([]);

  //structure data change
  const [sAnnotation, setAnnotationID] = useState(Number);
  const [sDateTime, setDateTime] = useState(String);
  const [sPharmacist, setPharmacistID] = useState(Number);
  const [sPrescription, setPrescriptionID] = useState(Number);

  const [sPharmacistName, setPharmacistName] = useState(String);
  const [sNamePatient, setNamePatient] = useState(String);
  const [sNameDoctor, setNameDoctor] = useState(String);

  //function get from api
  const getAnnotation = async () => {
    const res = await api.listAnnotation({ limit: 0, offset: 0 });
    setLoading(false);
    setAnnotation(res);
    console.log(res); ////////////////////
  };

  const getPrescription = async () => {
    const res = await api.listPrescription({ limit: 10, offset: 0 });
    setLoading(false);
    setPrescription(res);
    console.log(res); ////////////////////
  };

  const checkPosition = async () => {
    const PositionName = JSON.parse(
      String(localStorage.getItem('positiondata')),
    );
    setLoading(false);

    if (PositionName != 'pharmacist') {
      localStorage.clear();
      history.pushState('', '', './');
      window.location.reload(false);
    } else {
      setPharmacistID(Number(localStorage.getItem('pharmacist-id')));
      setPharmacistName(String(localStorage.getItem('pharmacist-name')));
    }
  };

  // Lifecycle Hooks
  useEffect(() => {
    getAnnotation();
    getPrescription();
    checkPosition();
  }, [loading]);

  const annotation_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setAnnotationID(event.target.value as number);
  };

  const datetime_handleChange = (event: any) => {
    setDateTime(event.target.value as string);
  };

  const CreateDispenseMedicines = async () => {
    if (
      sAnnotation != null &&
      sDateTime != null &&
      sPrescription != null &&
      sPharmacist != null
    ) {
      const dispensemedicines = {
        annotation: sAnnotation,
        datetime: sDateTime + ':00+07:00',
        pharmacist: sPharmacist,
        prescription: sPrescription,
      };
      console.log(dispensemedicines); ////////////////////

      const res: any = await api.createDispensemedicine({
        dispensemedicine: dispensemedicines,
      });
      setStatus(true);
      if (res.id != '') {
        setAlert(true);
        window.location.reload(false);
      }
    } else {
      setStatus(true);
      setAlert(false);
    }
  };

  function selectPrescriptions(id: any, namePatient: any, nameDoctor: any) {
    setPrescriptionID(Number(id));
    setNamePatient(String(namePatient));
    setNameDoctor(String(nameDoctor));
  }

  return (
    <Page theme={pageTheme.service}>
      <Header title={`${profile.givenName}`}>
        <Avatar src="/broken-image.jpg" />
        <div style={{ marginLeft: 10, color: 'white' }}>{sPharmacistName}</div>
      </Header>

      <Content>
        <ContentHeader title="">
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">บันทึกเรียบร้อย!</Alert>
              ) : (
                <Alert severity="warning" style={{ marginTop: 20 }}>
                  บันทึกไม่สำเร็จ!
                </Alert>
              )}
            </div>
          ) : null}
        </ContentHeader>
        <Grid
          container
          direction="row"
          justify="space-evenly"
          alignItems="stretch"
        >
          <Grid item xs>
            <TableContainer component={Paper}>
              <Table aria-label="simple table">
                <TableHead>
                  <TableRow>
                    <TableCell align="center">No.</TableCell>
                    <TableCell align="center">คนไข้</TableCell>
                    <TableCell align="center">เลือก</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {prescriptions.map((item: any) => (
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
                          Select
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
                      value={sPrescription}
                      inputProps={{ readOnly: true }}
                    ></TextField>
                  </FormControl>

                  <FormControl
                    variant="outlined"
                    className={classes.formControl}
                    fullWidth
                  >
                    <Typography className={classes.headLabel}>
                      ชื่อคนไข้
                    </Typography>
                    <TextField
                      label="ชื่อ-นามสกุล"
                      name="name"
                      variant="outlined"
                      value={sNamePatient}
                      inputProps={{ readOnly: true }}
                    ></TextField>
                  </FormControl>

                  <FormControl
                    variant="outlined"
                    className={classes.formControl}
                    fullWidth
                  >
                    <Typography className={classes.headLabel}>
                      ชื่อแพทย์ผู้ตรวจ
                    </Typography>
                    <TextField
                      label="ชื่อ-นามสกุล"
                      name="physician"
                      variant="outlined"
                      value={sNameDoctor}
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
                          <TableCell align="center">Name</TableCell>
                          <TableCell align="center">Brand</TableCell>
                          <TableCell align="center">Value</TableCell>
                        </TableRow>
                      </TableHead>
                      <TableBody>
                        {prescriptions.map((item: any) => (
                          <TableRow key={item.id}>
                            <TableCell align="center">
                              {item.edges?.prescriptionmedicine?.id}
                            </TableCell>
                            <TableCell align="center">
                              {item.edges?.prescriptionmedicine?.name}
                            </TableCell>
                            <TableCell align="center">
                              {item.edges?.prescriptionmedicine?.brand}
                            </TableCell>
                            <TableCell align="center">
                              {item.edges?.prescriptionmedicine?.amount}
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
                      id="date"
                      label="วันที่"
                      type="datetime-local"
                      value={sDateTime}
                      onChange={datetime_handleChange}
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
                    <TextField
                      select
                      label="เลือกหมายเหตุ"
                      name="annotation"
                      value={sAnnotation}
                      onChange={annotation_handleChange}
                      variant="outlined"
                    >
                      {annotations.map(option => (
                        <MenuItem key={option.id} value={option.id}>
                          {option.messages}
                        </MenuItem>
                      ))}
                    </TextField>
                  </FormControl>

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

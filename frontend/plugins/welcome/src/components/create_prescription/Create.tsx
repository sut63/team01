import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';

import {
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Link,
  Button,
} from '@material-ui/core';
import Timer from '../Timer';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntPatientInfo } from '../../api/models/EntPatientInfo';
import { ControllersPrescription } from '../../api/models/ControllersPrescription';
import { EntDoctor } from '../../api/models/EntDoctor';
import { EntMedicine } from '../../api/models/EntMedicine';
//import { Alert } from '@material-ui/lab';
import { Link as RouterLink } from 'react-router-dom';
//import { EntPrescription } from '../../api';
import { Cookies } from 'react-cookie/cjs'; //cookie
import Swal from 'sweetalert2';

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
    display: 'flex',
    justifyContent: 'center',
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    margin: theme.spacing(3),
    width: 320,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: '25ch',
  },

  margin: {
    margin: theme.spacing(3),
  },
  input: {
    display: 'none',
  },
}));

const NewPatientright: FC<{}> = () => {
  const classes = useStyles();
  const profile = { givenName: ' ' };
  const cookies = new Cookies();
  const http = new DefaultApi();

  const [Prescription, setPrescription] = React.useState<
    Partial<ControllersPrescription>
  >({});
  const [PatientInfo, setPatientInfo] = React.useState<EntPatientInfo[]>([]);
  const [Doctor, setDoctor] = React.useState<EntDoctor[]>([]);
  const [Medicine, setMedicine] = React.useState<EntMedicine[]>([]);
  const [status, setStatus] = React.useState(false);
  const [alert, setAlert] = React.useState(true);
  const [name, setName] = React.useState(String);
  const [id, setID] = React.useState(Number);

  const getMedicine = async () => {
    const res = await http.listMedicine({ limit: 10, offset: 0 });
    setMedicine(res);
  };

  const getPatientInfo = async () => {
    const res = await http.listPatientInfo({ limit: 10, offset: 0 });
    setPatientInfo(res);
  };

  const getDoctor = async () => {
    const res = await http.listDoctor({ limit: 10, offset: 0 });
    setDoctor(res);
  };

   // alert setting
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
  const checkPosition = async () => {
    if(cookies.get('PositionData') != 'Prescription'){
      history.pushState('', '', '/' + cookies.get('PositionData'));
      window.location.reload(false); 
    }
  };

  // Lifecycle Hooks
  useEffect(() => {
    getMedicine();
    getPatientInfo();
    getDoctor();
    setName(cookies.get('Name'));
    checkPosition();
    setPrescription({ doctorID: Number(cookies.get('ID')) });
  }, []);

  // set data to object Patientright
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof NewPatientright;
    const { value } = event.target;
    //const validateValue = value as string
    setPrescription({ ...Prescription, [name]: value });
  };

 
 
  const CreatePrescription = async () => {
    const apiUrl = 'http://localhost:8080/api/v1/Prescription';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(Prescription),
    };

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        console.log(data.status);
        if (data.status == true) {
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: data.error,
          });
        }
      });
  };


  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`ระบบสั่งจ่ายยา ${profile.givenName || 'to Backstage'}`}
       
      >
        <Timer />
      </Header>
      <Content>
        <ContentHeader title="ข้อมูล">
         
        </ContentHeader>

        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel>PatientInfo</InputLabel>
              <Select
                name="patientInfoID"
                value={Prescription.patientInfoID}
                onChange={handleChange}
              >
                {PatientInfo.map((item: any) => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.name}
                    </MenuItem>
                  );
                })}
              </Select>
            </FormControl>
          </form>
        </div>

        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl variant="outlined" className={classes.formControl}>
              <TextField
                name="doctor"
                label="doctor"
                value={name}
                InputProps={{
                  readOnly: true,
                }}
                variant="outlined"
              />
            </FormControl>
          </form>
        </div>

        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel>Medicine</InputLabel>
              <Select
                name="medicineID"
                value={Prescription.medicineID}
                onChange={handleChange}
              >
                {Medicine.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.name}
                    </MenuItem>
                  );
                })}
              </Select>
            </FormControl>
          </form>
        </div>

        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl variant="filled" className={classes.formControl}>
              <TextField
                name="value"
                label="Value"
                variant="outlined"
                type="string"
                size="medium"
                value={Prescription.value}
                onChange={handleChange}
              />
            </FormControl>
          </form>
        </div>

        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl variant="filled" className={classes.formControl}>
              <TextField
                name="symptom"
                label="Symptom อาการ"
                variant="outlined" 
                size="medium"
                value={Prescription.symptom}
                onChange={handleChange}
              />
            </FormControl>
          </form>
        </div>

        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl variant="filled" className={classes.formControl}>
              <TextField
                name="annotation"
                label="Annotation หมายเหตุ"
                variant="outlined"
                size="medium"
                value={Prescription.annotation}
                onChange={handleChange}
              />
            </FormControl>
          </form>
        </div>

        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <div className={classes.margin}>
              <Button
                onClick={() => {
                  CreatePrescription();
                }}
                variant="contained"
                color="primary"
              >
                Submit
              </Button>

              <Link component={RouterLink} to="/">
                <Button variant="contained" color="primary">
                  กลับสู่หน้าหลัก
                </Button>
              </Link>
            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
};
export default NewPatientright; 
import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme, ContentHeader, } from '@backstage/core';

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
import { EntDoctor } from '../../api/models/EntDoctor';
import { EntMedicine } from '../../api/models/EntMedicine';
import { Alert } from '@material-ui/lab';




import { Link as RouterLink } from 'react-router-dom';
import { EntPrescription } from '../../api';

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

interface Prescription_Type {
  /**
   * 
   * @type {number}
   * @memberof ControllersPrescription
   */
  doctorID?: number;
  /**
   * 
   * @type {number}
   * @memberof ControllersPrescription
   */
  medicineID?: number;
  /**
   * 
   * @type {number}
   * @memberof ControllersPrescription
   */
  patientInfoID?: number;
  /**
   * 
   * @type {string}
   * @memberof ControllersPrescription
   */
  value?: string;
}


const NewPatientright: FC<{}> = () => {
  const classes = useStyles();
  const profile = { givenName: 'สิทธ์' };
  const http = new DefaultApi();

  const [Prescription, setPrescription] = React.useState<Partial<Prescription_Type>>({});

  
  const [PatientInfo, setPatientInfo] = React.useState<EntPatientInfo[]>([]);

  const [Doctor, setDoctor] = React.useState<EntDoctor[]>([]);
  const [Medicine, setMedicine] = React.useState<EntMedicine[]>([]);

  const [status, setStatus] = React.useState(false);
  const [alert, setAlert] = React.useState(true);


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

  // Lifecycle Hooks
  useEffect(() => {
    getMedicine();
    getPatientInfo();
    getDoctor();
  }, []);

  // set data to object Patientright
  const handleChange = (

    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof NewPatientright;
    const { value } = event.target;
    setPrescription({ ...Prescription, [name]: value });
  };





  const CreatePatientright = async () => {

    const res: any = await http.createPrescription({
      prescription: Prescription


    });
    console.log(Prescription);
    setStatus(true);


    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }

    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);


  };



  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`ลงทะเบียน ${profile.givenName || 'to Backstage'}`}
        subtitle="Some quick intro and links."
      >
        <Timer />

      </Header>
      <Content>
        <ContentHeader title="ข้อมูล">

          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  บันทึกสำเร็จ
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    This is a warning alert — check it out!
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>

        <div>
          <br /><p>json {JSON.stringify(Prescription)} </p>
        </div>


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
              <InputLabel>Doctor</InputLabel>
              <Select
                name="doctorID"
                value={Prescription.doctorID}
                onChange={handleChange}
              >
                {Doctor.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.name}
                    </MenuItem>
                  );
                })}
              </Select>
            </FormControl>




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

        <ContentHeader title="พนักงาน" />





        <div className={classes.root}>
          <form noValidate autoComplete="off">

          <FormControl variant="filled" className={classes.formControl}>
                                <TextField
                                    name="value"
                                    label="Hight(cm)"
                                    variant="outlined"
                                    type="number"
                                    size="medium"
                                        
                                    value={Prescription.value}
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
                  CreatePatientright();
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
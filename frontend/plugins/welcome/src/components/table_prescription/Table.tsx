import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from './Tables';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
//import { EntUser } from '../../api/models/EntUser';
import { Cookies } from 'react-cookie/cjs';//cookie
import { useEffect } from 'react';
import { Alert } from '@material-ui/lab';
import { EntPatientInfo } from '../../api/models/EntPatientInfo';
import { Content, Header, Page, pageTheme,ContentHeader,Link,} from '@backstage/core';
import {FormControl,TextField} from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';
import Autocomplete from '@material-ui/lab/Autocomplete';

const Table: FC<{}> = () => {
  const profile = { givenName: 'ระบบค้นหารายการการสั่งจ่ายยา' };
  const http = new DefaultApi();
  const useStyles = makeStyles(theme => ({
    table: {
      minWidth: 650,
    },
    formControl: {
      margin: theme.spacing(3),
      width: 350,
    },
  }));
const [Sc, setSc] = React.useState<string>();
const [Pat, setPat] = React.useState<string>();
const classes = useStyles();
const [loading, setLoading] = React.useState(true);
const [status, setStatus] =   React.useState(false);
const [alert, setAlert] =  React.useState(true);
const [PatientInfo, setPatientInfo] = React.useState<EntPatientInfo[]>([]);
const getPatientInfo = async () => {
  const res = await http.listPatientInfo({ limit: 110, offset: 0 });
  setPatientInfo(res);
};
const handleChange = (event : any, value : unknown) => {
  setPat(value as string);
};
const So = async ()=>{
  setSc(Pat);
  var p = await http.listPatientInfo({ limit: 100, offset: 0 })
  let i = 0
  for (let u of p){
    if( u.name === Pat && u.edges?.patientprescription !== undefined)
    i = i+1
  console.log(i)  
  }
  //console.log("ผู้ป่วย = ",Pat)
  //console.log("p = ",p)

  if (i != 0){
   setStatus(true);
   setAlert(true);
  } else {
   setStatus(true);
   setAlert(false);
  }
   setTimeout(() => {
   setStatus(false);
  }, 5000);
 
};

 useEffect(() => {
    getPatientInfo();
    setLoading(false);
  }, [loading]);

   return (
    <Page theme={pageTheme.home}>
      <Header
        title={`ยินดีต้อนรับเข้าสู่${profile.givenName || 'to Backstage'}`}
       >
      </Header>
      <Content>
        <ContentHeader title="รายการการสั่งจ่ายยา">
        {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 พบรายการการสั่งจ่ายยา
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 ไม่พบรายการการสั่งจ่ายยา
               </Alert>
             )}
           </div>
         ) : null}


          <FormControl variant="outlined" className={classes.formControl}>
            <Autocomplete
              id="patientname"  
              options={PatientInfo.map((option) => option.name)}
              onChange={handleChange}
              renderInput={(params) => (
                <TextField {...params} label="ชื่อผู้ป่วย" margin="normal" variant="outlined" />
              )}
            />
            </FormControl>
          <Button
            onClick={() => {
              So();
            }}
            style={{ marginLeft: 10 }}
            variant="contained"
            color="primary" 
          >ค้นหา</Button>&emsp;
         
         <Link component={RouterLink} to="/">
            <Button variant="contained" color="primary" style={{ backgroundColor: "#FFcc00" }}> Home</Button>
          </Link>&emsp;

         <Link component={RouterLink} to="/Prescription">
            <Button variant="contained" color="primary"  style={{ backgroundColor: "#FF6666" }}>
              เพิ่มรายการ
           </Button>
          </Link>


        </ContentHeader>

        <ComponanceTable sim={Sc}></ComponanceTable>


      </Content>
    </Page>
  );
};

export default Table;
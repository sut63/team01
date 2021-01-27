import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from './Tables';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
//import { EntUser } from '../../api/models/EntUser';
import { Cookies } from 'react-cookie/cjs';//cookie
import { useEffect } from 'react';

import { EntPatientInfo } from '../../api/models/EntPatientInfo';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  Link,
} from '@backstage/core';
import {
  FormControl,
  Select,
  InputLabel,
  MenuItem,
} from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';

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
  const [Sc, setSc] = React.useState<number>(0);
const [Pat, setPat] = React.useState<number>(0);
const classes = useStyles();
const [loading, setLoading] = React.useState(true);
//const [Users, setUsers] = React.useState<Partial<EntUser>>();

const [PatientInfo, setPatientInfo] = React.useState<EntPatientInfo[]>([]);
const getPatientInfo = async () => {
    const res = await http.listPatientInfo({ limit: 110, offset: 0 });
    setPatientInfo(res);
  };
  const handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPat(event.target.value as number);
  };
  const So =()=>{
setSc(Pat)
  }
  useEffect(() => {
    getPatientInfo();
    setLoading(false);
  }, [loading]);

   return (
    <Page theme={pageTheme.home}>
      <Header
        title={`ยินดีต้อนรับ เข้าสู่ ${profile.givenName || 'to Backstage'}`}
       >
      </Header>
      <Content>
        <ContentHeader title="รายการการสั่งจ่ายยา">
          <FormControl variant="outlined" className={classes.formControl}>
            <InputLabel>ผู้ป่วย</InputLabel>
            <Select
              name="patientinfo"
              value={Pat}
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
          <Button
            onClick={() => {
              So();
            }}
            style={{ marginLeft: 10 }}
            variant="contained"
            color="primary" 
          
          >
            ค้นหา
               </Button>&emsp;
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
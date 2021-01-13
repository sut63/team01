import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  ContentHeader,
  Header,
  Page,
  pageTheme,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';

import { DefaultApi } from '../../api/apis';
import { InputLabel, MenuItem, Select } from '@material-ui/core';
import Typography from '@material-ui/core/Typography';
import TableCell from '@material-ui/core/TableCell';
import Avatar from '@material-ui/core/Avatar';
import Box from '@material-ui/core/Box';
import { Alert } from '@material-ui/lab';


import { EntMedicine } from '../../api/models/EntMedicine';
import { EntCompany } from '../../api/models/EntCompany';
import { EntPharmacist } from '../../api/models/EntPharmacist';

import Swal from 'sweetalert2';
import SaveIcon from '@material-ui/icons/Save'; // icon save

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(1),
    },
    withoutLabel: {
      marginTop: theme.spacing(2),
    },
    textField: {
      width: '25ch',
    },
  }),
);

const Toast = Swal.mixin({
  toast: true,
  position: 'top-end',
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
  didOpen: toast => {
    toast.addEventListener('mouseenter', Swal.stopTimer);
    toast.addEventListener('mouseleave', Swal.resumeTimer);
  },
});

interface oorder {
  pharmacistid: number;
  medicineid: number;
  companyid: number;
  price: number;
  amount: number;
  addedtime: Date;
}

export default function Order() {
  const classes = useStyles();
  const profile = { givenName: 'to Order ' };
  const api = new DefaultApi();
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);

  const [oorder, setOreder] = React.useState<Partial<oorder>>({});

  const [medicines, setMedicines] = React.useState<EntMedicine[]>([]);
  const [companys, setCompanys] = React.useState<EntCompany[]>([]);
  const [pharmacists, setPharmacists] = React.useState<EntPharmacist[]>([]);

  const [pharmacistID, setPharmacistid] = useState(Number);
  const [medicineID, setMedicineid] = useState(Number);
  const [companyID, setCompanyid] = useState(Number);
  const [orderpriceid, setOrderpriceid] = useState(Number);
  const [orderamountid, setOrderamountid] = useState(Number);
  const [datetime, setDatetime] = useState(String);


  let price = Number(orderpriceid)
  let amount = Number(orderamountid)
  let pharmacistid = Number(pharmacistID)
  let medicineid = Number(medicineID)
  let companyid = Number(companyID)
  
  console.log(pharmacistID)
  useEffect(() => {

    const getpharmacists = async () => {

      const mn = await api.listPharmacist({ limit: 10, offset: 0 });
      setLoading(false);
      setPharmacists(mn);
    };
    getpharmacists();



    const getmedicines = async () => {

      const pr = await api.listMedicine({ limit: 10, offset: 0 });
      setLoading(false);
      setMedicines(pr);
    };
    getmedicines();

    const getcompanys = async () => {

      const pay = await api.listCompany({ limit: 10, offset: 0 });
      setLoading(false);
      setCompanys(pay);
    };
    getcompanys();

  }, [loading]);

  const order = {
    pharmacistid,
    medicineid,
    companyid,
    amount,
    price,
    addedtime: datetime + ":00+07:00"
  }
  console.log(order)

  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/orders';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(order),
    };

    console.log(order); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.id != null) {
          //clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }

  const Pharmacist_id_handleChange = (event: any) => {
    setPharmacistid(event.target.value);
  };

  const Medicine_id_handleChange = (event: any) => {
    setMedicineid(event.target.value);
  }
  const Company_id_handleChange = (event: any) => {
    setCompanyid(event.target.value);
  };

  const Orderprice_id_handleChange = (event: any) => {
    setOrderpriceid(event.target.value);
  };
  const Orderamount_id_handleChange = (event: any) => {
    setOrderamountid(event.target.value);
  };
  const handleDatetimeChange = (event: any) => {
    setDatetime(event.target.value as string);
  };

 
  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`Welcome ${profile.givenName || 'to MedicineOrder'}`}
        subtitle="เลือกยาที่สั่งซื้อ."
      >

        <Avatar>C</Avatar>
        <Typography component="div" variant="body1">
          <Box color=""> เภสัชกร</Box>
          <Box color="secondary.main"></Box>
        </Typography>

      </Header>
      <Content>
      <ContentHeader title=" ">          
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  This is a success alert — check it out!
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    This is a warning alert — check it out!
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>

        <div className={classes.root}>
          <form noValidate autoComplete="off">



            <TableCell align="left">



              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 600 }}

              >
                <InputLabel id="medicine_id-label">Medicine</InputLabel>
                <Select
                  labelId="medicine_id-label"
                  label="Medicine"
                  id="medicine_id"
                  value={medicineID || ''}
                  onChange={Medicine_id_handleChange}
                  style={{ width: 300 }}
                >
                  {medicines.map((item: EntMedicine) =>
                    <MenuItem value={item.id}>{item.name}</MenuItem>)}
                </Select>
              </FormControl>

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 600 }}
              >
                <InputLabel id="company_id-label">Company</InputLabel>
                <Select
                  labelId="company_id-label"
                  label="Company"
                  id="company_id"
                  value={companyID || ''}
                  onChange={Company_id_handleChange}
                  style={{ width: 300 }}
                >
                  {companys.map((item: EntCompany) =>
                    <MenuItem value={item.id}>{item.name}</MenuItem>)}
                </Select>
              </FormControl>

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 600 }}
              >
                <InputLabel id="pharmacist_id-label">Pharmacist</InputLabel>
                <Select
                  labelId="pharmacist_id-label"
                  label="Pharmacist"
                  id="pharmacist_id"
                  value={pharmacistID || ''}
                  onChange={Pharmacist_id_handleChange}
                  style={{ width: 300 }}
                >
                  {pharmacists.map((item: EntPharmacist) =>
                    <MenuItem value={item.id}>{item.name}</MenuItem>)}
                </Select>
              </FormControl>


              
              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 302 }}
              >
                <TextField id="outlined-number" type='number' InputLabelProps={{
                  shrink: true,
                }} label="จำนวนที่สี่ง" variant="outlined"
                  onChange={Orderamount_id_handleChange}
                  value={orderamountid || ''}
                />
              </FormControl>

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 302 }}
              >
         
                
              </FormControl>
              

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 302 }}
              >
                <TextField id="outlined-number" type='number' InputLabelProps={{
                  shrink: true,
                }} label="ราคาสุทธิ" variant="outlined"
                  onChange={Orderprice_id_handleChange}
                  value={orderpriceid || ''}
                />
              </FormControl>




              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginLeft: 560, width: 600 }}
              >
                <TextField
                  id="date"
                  label="วัน-เวลาที่สั่งซื้อ"
                  type="datetime-local"
                  value={datetime || ''}
                  onChange={handleDatetimeChange}
                  //defaultValue="2020-05-24"
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </FormControl>




            </TableCell>

            <div className={classes.margin}>
            <TableCell align="right">
                <Button
                component={RouterLink}
                to="/Order"
                  variant="contained"
                  color="primary"
                  size="large"
                  style={{ marginLeft: 545, width: 200 }}
                  className={classes.margin}
                  onClick={() => {
                    save();
                  }}
                  startIcon={<SaveIcon
                  />}
                >
                  Save
                </Button>
              </TableCell>

             

             

            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
}
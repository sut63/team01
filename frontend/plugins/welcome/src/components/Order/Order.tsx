import React, { useState, useEffect,FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { Cookies } from 'react-cookie/cjs'; //cookie
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
import { AppBar, InputLabel, MenuItem, Select } from '@material-ui/core';
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


  const oorder: FC<{}> = () => {
  const classes = useStyles();
  const profile = { givenName: 'to Order ' };
  const api = new DefaultApi();
  const cookies = new Cookies();
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);

  const [oorDer, setOreder] = React.useState<Partial<oorder>>({});
  

  const [medicines, setMedicines] = React.useState<EntMedicine[]>([]);
  const [companys, setCompanys] = React.useState<EntCompany[]>([]);
 

  const [pharmacistID, setPharmacistid] = React.useState(String);
  const [medicineID, setMedicineid] = useState(Number);
  const [companyID, setCompanyid] = useState(Number);
  const [orderpriceid, setOrderpriceid] = useState(Number);
  const [hospitalid, setOrderhospitalid] = useState(String);
  const [orderamountid, setOrderamountid] = useState(Number);
  const [datetime, setDatetime] = useState(String);

  const [orderpriceidError, SetPriceIDError] = React.useState('');
  const [orderhospitalidError, SetHospitalIDError] = React.useState('');
  const [orderamountidError, SetAmountIDError] = React.useState('');
  const [errors, setError] = React.useState(String);


  let price = Number(orderpriceid)
  let hospitalId = String(hospitalid)
  let amount = Number(orderamountid)
  let medicineid = Number(medicineID)
  let companyid = Number(companyID)
  
  console.log(pharmacistID)
  useEffect(() => {

  



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

  const handleChange = (event: React.ChangeEvent<{ name?: string; value: unknown }>) => {
    const name = event.target.name as keyof typeof oorder;
    const { value } = event.target;
    //const validateValue = value.toString()
    //checkPattern(name, validateValue)
    setOreder({ ...oorDer, [name]: value });
    //console.log(oorder);
};

/*const validatePhoneID = (val: string) => {
return val.length == 9 ? true : false;
}

const validateAmountID = (val: number) => {
  return val > 1 ;
  }

  const validatePriceID = (val: number) => {
    return val > 1 ;
    }



const checkPattern = (id:number ,value: number ) => {
  switch(id) {
    case 'amount':
      validateAmountID(value) ? SetAmountIDError('') :  SetAmountIDError('กรอกจำนวนยา');
      return;
    case 'price':
      validatePriceID(value) ? SetPriceIDError('') :  SetPriceIDError('กรอกราคาสุทธิ');
      return;
   case 'phone':
      validatePhoneID(value) ? SetPhoneIDError('') :  SetPhoneIDError('กรอกเบอร์โทรศัพท์ 9 ตัว');
      return;
  }
}*/



const alertMessage = (icon: any , title: any) => {
  Toast.fire({
    icon: icon,
    title: title,
  })
}



const checkCaseSaveError = (field : String) => {
  switch(field) {
    case 'amount' :
      alertMessage("error","กรอกจำนวนสินค้าไม่ถูกต้อง");
      return;
    case 'price' :
      alertMessage("error","กรอกราคาสินค้าไม่ถุกต้อง");
      return;
      case 'hospitalid' :
        alertMessage("error","รูปแบบหมายเลขติดต่อหน่วยงานไม่ถุกต้อง");
        return;
    default:
      alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
      return;
  }
}

  //console.log(order)

  function save() {

    const order = {
      pharmacistid: Number(oorDer.pharmacistid),
      medicineid,
      companyid,
      amount,
      hospitalId,
      price,
      addedtime: datetime + ":00+07:00"
    };
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
        if (data.status == true) {
          //clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          checkCaseSaveError(data.error.Name)
        }
      });
  }



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
  const Orderhospital_id_handleChange = (event: any) => {
    setOrderhospitalid(event.target.value);
  };
  const handleDatetimeChange = (event: any) => {
    setDatetime(event.target.value as string);
  };

  const clear = async () => {
    setOreder({});
    setOreder({ pharmacistid: Number(cookies.get('ID')) });
}
  useEffect(() => {
  
    setPharmacistid(cookies.get('Name'));
    setOreder({ pharmacistid: Number(cookies.get('ID')) });
    
}, []);

 
  return (
    <Page theme={pageTheme.home}>
      <Header
        
        title={`Welcome to MedicineOrder `}
        subtitle="เลือกยาที่สั่งซื้อ"
      >
     

        <Avatar></Avatar>
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



            <FormControl variant='outlined' style={{ marginLeft: 693, width: 302 }}>
                         <InputLabel id="pharmacist"></InputLabel>
                         <div className={classes.margin}>
                                    <div>
                                        <TextField
                                            style={{ width: 302 }}
                                            name="pharmacist"
                                            label="เภสัชกร"
                                            value={pharmacistID}
                                            InputProps={{
                                                readOnly: true,
                                            }}
                                            variant="outlined"
                                        />
                                    </div></div>
                                    </FormControl>

                                     <FormControl variant='outlined' style={{ marginTop: 20 , marginLeft: 700, width: 302 }}>
                                        <InputLabel id="medicine">ยา</InputLabel>
                                        <Select
                                            labelId="medicine"
                                            name="medicine"
                                            onChange={Medicine_id_handleChange}
                                            label='ยา'
                                        >
                                            {medicines.map(item => {
                                                return (
                                                    <MenuItem key={item.id} value={item.id}>
                                                        {item.name}
                                                    </MenuItem>
                                                );
                                            })}
                                        </Select>
                                    </FormControl>
                                    




              
              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{marginTop: 20 , marginLeft: 700, width: 302 }}
              >
                <TextField id="amount" type='number' InputLabelProps={{
                  shrink: true,
                }} label="จำนวนสินค้า" variant="outlined"
                  onChange={Orderamount_id_handleChange}
                  value={orderamountid || ''}
                />
              </FormControl>

 
              

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{ marginTop: 20 ,marginLeft: 700, width: 302 }}
              >
                <TextField id="price" type='number' InputLabelProps={{
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
                style={{ marginTop: 20 ,marginLeft: 700, width: 302 }}
              >
                <TextField
                  id="hospitalid"
                  label="เบอร์ติดต่อหน่วยงาน"
                  variant="outlined"
                  type="string"
                  size="medium"
                  value={hospitalid}
                  onChange={Orderhospital_id_handleChange}
                  style={{ marginRight: 300, width: 300 }}
                />
              </FormControl>

  


              
              <FormControl variant='outlined' style={{ marginTop: 20 , marginLeft: 700, width: 302 }}>
                                        <InputLabel id="company">บริษัท</InputLabel>
                                        <Select
                                            labelId="company"
                                            name="company"
                                            onChange={Company_id_handleChange}
                                            label='บริษัท'
                                        >
                                            {companys.map(item => {
                                                return (
                                                    <MenuItem key={item.id} value={item.id}>
                                                        {item.name}
                                                    </MenuItem>
                                                );
                                            })}
                                        </Select>
                                    </FormControl>




              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
                style={{marginTop: 20 , marginLeft: 750, width: 600 }}
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
                <Button
                    style={{ marginLeft: 40 }}
                    component={RouterLink}
                    to="/SearchOrder"
                    variant="contained"
                    color="secondary"
                  >
                    ค้นหาข้อมูลสั่งซื้อยาย้อนหลัง
                </Button>
              </TableCell>

             

             

            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
 
};

export default oorder;
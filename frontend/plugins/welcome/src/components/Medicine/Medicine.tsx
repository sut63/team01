import React, {FC, useEffect, useState} from 'react';
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
  Grid,
  Card,
  CardContent,
} from '@material-ui/core';
import SendIcon from '@material-ui/icons/Send';
import Swal from 'sweetalert2'; // alert

import { DefaultApi } from '../../api/apis/'; // Api Gennerate From Command
import { EntMedicineType } from '../../api/models/EntMedicineType'; // import interface MedicineType
import { EntLevelOfDangerous } from '../../api/models/EntLevelOfDangerous'; // import interface LevelOfDangerous
import { EntUnitOfMedicine } from '../../api/models/EntUnitOfMedicine'; // import interface UnitOfMedicine

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    paper: {
      marginTop: theme.spacing(2),
      marginBottom: theme.spacing(2),
      marginLeft: theme.spacing(2),
    },
    formControl: {
      width: 500,
    },
    textField: {
      width: 400,
    },
    comboBox: {
      width: 400,
    },
    cardBox: {
      width: 1000,
    },
    button: {
      margin: theme.spacing(1),
      height: 45,
      marginLeft: 20,
    },
    large: {
      width: theme.spacing(6),
      height: theme.spacing(6),
    },
  }),
);


const Medicine: FC<{}> = () => {
  const classes = useStyles();
  const api = new DefaultApi();
  const profile = { givenName: 'ระบบบันทึกข้อมูลยา' };
  const [status, setStatus] = React.useState(false);
  

  const [medicinetype, setMedicineType] = React.useState<EntMedicineType[]>([]);
  const [levelOfdangerous, setLevelOfDangerous] = React.useState<EntLevelOfDangerous[]>([]);
  const [unitofmedicine, setUnitOfMedicine] = React.useState<EntUnitOfMedicine[]>([]);

  const [sMName, setMName] = React.useState(String);
  const [sMSerial, setMSerial] = React.useState(String);
  const [sMBrand, setMBrand] = React.useState(String);
  const [sMAmount, setMAmount] = React.useState(Number);
  const [sMPrice, setMPrice] = React.useState(Number);
  const [sMHowtouse, setMHowtouse] = React.useState(String);
  const [sMLevelOfDangerous, setMLevelOfDangerous] = React.useState(Number);
  const [sMMedicineType, setMMedicineType] = React.useState(Number);
  const [sMUnitOfMedicine, setMUnitOfMedicine] = React.useState(Number);

  const sMName_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setMName(event.target.value as string);
  };

  const sMSerial_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setMSerial(event.target.value as string);
  };

  const sMBrand_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setMBrand(event.target.value as string);
  };

  const sMAmount_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setMAmount(event.target.value as number);
  };

  const sMPrice_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setMPrice(event.target.value as number);
  };

  const sMHowtouse_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setMHowtouse(event.target.value as string);
  };

  const sMLevelOfDangerous_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setMLevelOfDangerous(event.target.value as number);
  };

  const sMMedicineType_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setMMedicineType(event.target.value as number);
  };

  const sMUnitOfMedicine_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setMUnitOfMedicine(event.target.value as number);
  };



  const getMedicineType = async () => {
    const res = await api.listMedicineType({ limit: 10, offset: 0 });
    setMedicineType(res);

  };

  const getLevelOfDangerous = async () => {
    const res = await api.listLevelOfDangerous({ limit: 10, offset: 0 });
    setLevelOfDangerous(res);

  };

  const getUnitOfMedicine = async () => {
    const res = await api.listUnitOfMedicine({ limit: 10, offset: 0 });
    setUnitOfMedicine(res);
 
  };

  // Lifecycle Hooks
  useEffect(() => {
    getMedicineType();
    getLevelOfDangerous();
    getUnitOfMedicine();
  }, []);

  // alert setting
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

  const CreateMedicine = async () => {
    const smedicine = {
      levelOfDangerousID: sMLevelOfDangerous,
      medicineTypeID: sMMedicineType,
      unitOfMedicineID: sMUnitOfMedicine,
      Name: sMName,
      Serial: sMSerial,
      Brand: sMBrand,
      Amount: Number(sMAmount),
      Price: Number(sMPrice),
      Howtouse: sMHowtouse,
    };
    console.log(smedicine);
    const res: any = await api.createMedicine({ medicine: smedicine });
    setStatus(true);
    if (res.id != '' || res.id == undefined) {
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

    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };

  

  return (
    <Page theme={pageTheme.service}>
      <Header title={'ระบบบันทึกข้อมูลยา'}>
        
      </Header>

      <Content>
        <div className={classes.root}>
          <Card className={classes.cardBox}>
            <CardContent>
              <Container maxWidth="sm">
                <div style={{ textAlign: 'center' }}>
                  <ContentHeader title="การบันทึกข้อมูลยาในฐานระบบห้องยา"></ContentHeader>
                </div>
                <Grid container spacing={3}>
                  <Grid item xs={12}></Grid>
                  <Grid item xs={3}>
                    <div className={classes.paper}>Medicine Name</div>
                  </Grid>
                  <Grid item xs={9}>
                    <FormControl
                      variant="outlined"
                      className={classes.formControl}
                    >
                      <TextField
                        id="mname"
                        label="กรุณาระบุชื่อของยา"
                        type="string"
                        variant="outlined"
                        value={sMName}
                        onChange={sMName_handleChange}
                        className={classes.textField}
                      />
                    </FormControl>
                  </Grid>

                  <Grid item xs={3}>
                    <div className={classes.paper}>Serial Name</div>
                  </Grid>
                  <Grid item xs={9}>
                    <FormControl
                      variant="outlined"
                      className={classes.formControl}
                    >
                      <TextField
                        id="mserial"
                        label="กรุณาระบุรหัสยา"
                        type="string"
                        variant="outlined"
                        value={sMSerial}
                        onChange={sMSerial_handleChange}
                        className={classes.textField}
                      />
                    </FormControl>
                  </Grid>

                  <Grid item xs={3}>
                    <div className={classes.paper}>Brand Name</div>
                  </Grid>
                  <Grid item xs={9}>
                    <FormControl
                      variant="outlined"
                      className={classes.formControl}
                    >
                      <TextField
                        id="mbrand"
                        label="กรุณาระบุยี่ห้อยา"
                        type="string"
                        variant="outlined"
                        value={sMBrand}
                        onChange={sMBrand_handleChange}
                        className={classes.textField}
                      />
                    </FormControl>
                  </Grid>

                  <Grid item xs={3}>
                    <div className={classes.paper}>Stock Amount</div>
                  </Grid>
                  <Grid item xs={9}>
                    <FormControl
                      variant="outlined"
                      className={classes.formControl}
                    >
                      <TextField
                        id="mamount"
                        label="กรุณาระบุจำนวนยา"
                        type="number"
                        variant="outlined"
                        value={sMAmount}
                        onChange={sMAmount_handleChange}
                        className={classes.textField}
                      />
                    </FormControl>
                  </Grid>

                  <Grid item xs={3}>
                    <div className={classes.paper}>Medicine Price Per Unit</div>
                  </Grid>
                  <Grid item xs={9}>
                    <FormControl
                      variant="outlined"
                      className={classes.formControl}
                    >
                      <TextField
                        id="mprice"
                        label="กรุณาระบุราคายา ต่อ 1 หน่วยยา"
                        type="number"
                        variant="outlined"
                        value={sMPrice}
                        onChange={sMPrice_handleChange}
                        className={classes.textField}
                      />
                    </FormControl>
                  </Grid>

                  <Grid item xs={3}>
                    <div className={classes.paper}>Medicine Menaul</div>
                  </Grid>
                  <Grid item xs={9}>
                    <FormControl
                      variant="outlined"
                      className={classes.formControl}
                    >
                      <TextField
                        id="mhowtouse"
                        label="กรุณาระบุวิธีการใช้ยา"
                        type="string"
                        variant="outlined"
                        value={sMHowtouse}
                        onChange={sMHowtouse_handleChange}
                        className={classes.textField}
                      />
                    </FormControl>
                  </Grid>
        
                  <Grid item xs={3}>
                    <div className={classes.paper}>LevelOfDangerous</div>
                  </Grid>
                  <Grid item xs={9}>
                    <FormControl
                      variant="outlined"
                      className={classes.formControl}
                    >
                      <TextField
                        id="select levelofdangerous"
                        select
                        label="กรุณาระบุระดับความอันตรายของยา"
                        className={classes.textField}
                        value={sMLevelOfDangerous}
                        onChange={sMLevelOfDangerous_handleChange}
                        variant="outlined"
                      >
                        {levelOfdangerous.map(option => (
                          <MenuItem key={option.id} value={option.id}>
                            {option.name}
                          </MenuItem>
                        ))}
                      </TextField>
                    </FormControl>
                  </Grid>
                  
                  <Grid item xs={3}>
                    <div className={classes.paper}>MedicineType</div>
                  </Grid>
                  <Grid item xs={9}>
                    <FormControl
                      variant="outlined"
                      className={classes.formControl}
                    >
                      <TextField
                        id="select medicinetype"
                        select
                        label="กรุณาระบุประเภทของยา"
                        className={classes.textField}
                        value={sMMedicineType}
                        onChange={sMMedicineType_handleChange}
                        variant="outlined"
                      >
                        {medicinetype.map(option => (
                          <MenuItem key={option.id} value={option.id}>
                            {option.name}
                          </MenuItem>
                        ))}
                      </TextField>
                    </FormControl>
                  </Grid>

                  <Grid item xs={3}>
                    <div className={classes.paper}>UnitOfMedicine</div>
                  </Grid>
                  <Grid item xs={9}>
                    <FormControl
                      variant="outlined"
                      className={classes.formControl}
                    >
                      <TextField
                        id="select unitofmedicine"
                        select
                        label="กรุณาระบุหน่วยยา"
                        className={classes.textField}
                        value={sMUnitOfMedicine}
                        onChange={sMUnitOfMedicine_handleChange}
                        variant="outlined"
                      >
                        {unitofmedicine.map(option => (
                          <MenuItem key={option.id} value={option.id}>
                            {option.name}
                          </MenuItem>
                        ))}
                      </TextField>
                    </FormControl>
                  </Grid>

                  <Grid item xs={3}></Grid>
                  <Grid item xs={9}>
                    <Button
                      variant="contained"
                      color="primary"
                      className={classes.button}
                      startIcon={<SendIcon />}
                      onClick={() => {
                        CreateMedicine();
                      }}
                      
                    >
                      บันทึกการสมัคร
                    </Button>

                    <Button
                      variant="contained"
                      className={classes.button}
                      component={RouterLink}
                      to="/"
                    >
                      ย้อนกลับ
                    </Button>
                  </Grid>
                </Grid>
              </Container>
            </CardContent>
          </Card>
        </div>
      </Content>
    </Page>
  );
};


export default Medicine;
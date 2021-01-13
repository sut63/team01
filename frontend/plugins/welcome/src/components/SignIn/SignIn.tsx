import React, { FC, useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { makeStyles } from '@material-ui/core/styles';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import {
  Container,
  TextField,
  Button,
  Typography,
  Avatar,
  Card,
  CardContent,
  CssBaseline,
} from '@material-ui/core';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import VpnKeyOutlinedIcon from '@material-ui/icons/VpnKeyOutlined';
import { Alert } from '@material-ui/lab'; // alert
import { DefaultApi } from '../../api/apis';
import { EntPharmacist } from '../../api/models/EntPharmacist';
import { EntDoctor } from '../../api/models/EntDoctor';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
    display: 'flex',
    flexWrap: 'wrap',
    justifyContent: 'center',
  },
  paper: {
    marginTop: theme.spacing(2),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    width: theme.spacing(7),
    height: theme.spacing(7),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  cardBox: {
    marginTop: theme.spacing(7),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

const SignIn: FC<{}> = () => {
  const classes = useStyles();
  const profile = { givenName: 'ยินดีต้อนรับ' };
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(Boolean);

  const api = new DefaultApi();
  const [pharmacists, setPharmacist] = useState<EntPharmacist[]>([]);
  const [doctors, setDoctor] = useState<EntDoctor[]>([]);

  const [emails, setEmail] = React.useState(String);
  const [password, setPassword] = React.useState(String);

  const emailHandleChange = (event: any) => {
    setEmail(event.target.value as string);
  };
  const passwordHandleChange = (event: any) => {
    setPassword(event.target.value as string);
  };

  const getPharmacist = async () => {
    const res: any = await api.listPharmacist({ limit: 0, offset: 0 });
    setPharmacist(res);
  };

  const getDoctor = async () => {
    const res: any = await api.listDoctor({ limit: 0, offset: 0 });
    setDoctor(res);
  };

  const resetData = async () => {
    localStorage.setItem("pharmacist-id",JSON.stringify(null));
    localStorage.setItem("pharmacist-name",JSON.stringify(null));
    localStorage.setItem("positiondata",JSON.stringify(null));
    localStorage.setItem("doctor-id",JSON.stringify(null));
    localStorage.setItem("doctor-name",JSON.stringify(null));
  };

  useEffect(() => {
    getPharmacist();
    getDoctor();
    resetData();
  }, []);

  const SigninCheck = async () => {
    var checkPharmacist = false;

    pharmacists.map((item: any) => {
      console.log(item.email);
      if (item.email == emails && item.password == password) {
        setAlert(true);
        checkPharmacist = true;
        localStorage.setItem('pharmacist-id', JSON.stringify(item.id));
        localStorage.setItem('pharmacist-name', JSON.stringify(item.name));
        localStorage.setItem('positiondata', JSON.stringify('pharmacist'));
        history.pushState('', '', '/drug_allergy'); /////////มาแก้ที่หลังไปยังหน้ารวมเภสัชกร
        window.location.reload(false);
      }
    });

    if (checkPharmacist == false){
      doctors.map((item: any) => {
        console.log(item.email);
        if (item.email == emails && item.password == password) {
          setAlert(true);
          localStorage.setItem('doctor-id', JSON.stringify(item.id));
          localStorage.setItem('doctor-name', JSON.stringify(item.name));
          localStorage.setItem('positiondata', JSON.stringify('doctor'));
          history.pushState('', '', '/create_prescription');
          window.location.reload(false);
        }
      });
    }
    
    setStatus(true);
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };

  return (
    <Page theme={pageTheme.service}>
      <Header title={`${profile.givenName || 'to Backstage'}`}></Header>
      <Content>
        <ContentHeader title="ระบบห้องยา">
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">เข้าสู่ระบบสำเร็จ</Alert>
              ) : (
                <Alert severity="warning" style={{ marginTop: 20 }}>
                  ไม่พบข้อมูลในระบบ
                </Alert>
              )}
            </div>
          ) : null}
        </ContentHeader>
        <div className={classes.root}>
          <Card className={classes.cardBox}>
            <CardContent>
              <Container component="main" maxWidth="xs">
                <CssBaseline />
                <div className={classes.paper}>
                  <Avatar className={classes.avatar}>
                    <VpnKeyOutlinedIcon />
                    <LockOutlinedIcon />
                  </Avatar>
                  <Typography component="h1" variant="h5">
                    Sign in
                  </Typography>
                  <form className={classes.form} noValidate>
                    <TextField
                      variant="outlined"
                      margin="normal"
                      required
                      fullWidth
                      id="email"
                      label="Email Address"
                      name="email"
                      autoComplete="email"
                      onChange={emailHandleChange}
                      autoFocus
                    />

                    <TextField
                      variant="outlined"
                      margin="normal"
                      required
                      fullWidth
                      name="password"
                      label="Password"
                      type="password"
                      id="password"
                      onChange={passwordHandleChange}
                      autoComplete="current-password"
                    />

                    <Button
                      fullWidth
                      variant="contained"
                      color="primary"
                      className={classes.submit}
                      onClick={SigninCheck}
                    >
                      Sign In
                    </Button>
                  </form>
                </div>
              </Container>
            </CardContent>
          </Card>
        </div>
      </Content>
    </Page>
  );
};

export default SignIn;

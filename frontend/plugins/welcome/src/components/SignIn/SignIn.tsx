import React, { FC, useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
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
import { DefaultApi } from '../../api/apis';
import { EntPharmacist } from '../../api/models/EntPharmacist';
import { EntDoctor } from '../../api/models/EntDoctor';
import { Cookies } from 'react-cookie/cjs'; //cookie
import Swal from 'sweetalert2'; // alert

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
    marginTop: theme.spacing(2),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

const SignIn: FC<{}> = () => {
  const classes = useStyles();
  const profile = { givenName: 'ระบบห้องยา' };
  const cookies = new Cookies();
  const api = new DefaultApi();

  const [pharmacists, setPharmacist] = useState<EntPharmacist[]>([]);
  const [doctors, setDoctor] = useState<EntDoctor[]>([]);

  const [emails, setEmail] = React.useState(String);
  const [password, setPassword] = React.useState(String);
  const [emailError, setEmailError] = React.useState('');
  const [passwordError, setPasswordError] = React.useState('');

  // alert setting
  const Toast = Swal.mixin({
    //toast: true,
    position: 'center',
    showConfirmButton: false,
    timer: 1400,
    timerProgressBar: true,
  });

  const emailHandleChange = (event: any) => {
    setEmail(event.target.value as string);
    checkPattern('email', event.target.value as string);
  };
  const passwordHandleChange = (event: any) => {
    setPassword(event.target.value as string);
    checkPattern('password', event.target.value as string);
  };

  const getPharmacist = async () => {
    const res: any = await api.listPharmacist({ limit: 0, offset: 0 });
    setPharmacist(res);
  };

  const getDoctor = async () => {
    const res: any = await api.listDoctor({ limit: 0, offset: 0 });
    setDoctor(res);
  };

  const resetUserData = async () => {
    cookies.set('ID', JSON.stringify(0), { path: '/' });
    cookies.set('Name', JSON.stringify(null), { path: '/' });
    cookies.set('PositionData', JSON.stringify(''), { path: '/' });
    cookies.set('StatusLogin', JSON.stringify('No'), { path: '/' });
  };

  const checkPosition = async () => {
    if(cookies.get('StatusLogin') == 'Yes'){
      history.pushState('', '', '/' + cookies.get('PositionData'));
      window.location.reload(false);
    }else{
      resetUserData();
    }
  };

  useEffect(() => {
    getPharmacist();
    getDoctor();
    checkPosition();
  }, []);

  const validateEmail = (email: string) => {
    const re = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(email);
  };

  const validatePassword = (val: string) => {
    const regexpText = new RegExp('[`~!@#$%^&*_;?<>]');
    if (regexpText.test(val)) {
      return false;
    } else {
      return true;
    }
  };

  const checkPattern = (id: string, value: string) => {
    switch (id) {
      case 'email':
        validateEmail(value)
          ? setEmailError('')
          : setEmailError('รูปแบบอีเมลไม่ถูกต้อง');
        return;
      case 'password':
        validatePassword(value)
          ? setPasswordError('')
          : setPasswordError('รูปแบบอีเมลไม่ถูกต้อง ห้ามป้อนอักษรพิเศษ');
        return;
      default:
        return;
    }
  };

  const SigninCheck = async () => {
    var checkPharmacist = false;
    var showAlertFalse = true;

    pharmacists.map((item: EntPharmacist) => {
      if (item.email == emails && item.password == password) {
        checkPharmacist = true;
        showAlertFalse = false;
        cookies.set('ID', JSON.stringify(item.id), { path: '/' });
        cookies.set('Name', JSON.stringify(item.name), { path: '/' });
        cookies.set('StatusLogin', JSON.stringify('Yes'), { path: '/' });
        cookies.set(
          'PositionData',
          JSON.stringify(item.edges?.positioninpharmacist?.position),
          { path: '/' },
        );
      }
    });

    if (checkPharmacist == false) {
      doctors.map((item: any) => {
        if (item.email == emails && item.password == password) {
          showAlertFalse = false;
          cookies.set('ID', JSON.stringify(item.id), { path: '/' });
          cookies.set('Name', JSON.stringify(item.name), { path: '/' });
          cookies.set('StatusLogin', JSON.stringify('Yes'), { path: '/' });
          cookies.set('PositionData', JSON.stringify('Prescription'), {
            path: '/',
          });
        }
      });
    }

    if (showAlertFalse === true) {
      Toast.fire({
        icon: 'error',
        title: 'เข้าสู่ระบบไม่สำเร็จ กรุณาตรวจสอบ Email หรือ Password',
      });
    } else {
      Toast.fire({
        icon: 'success',
        title: 'เข้าสู่ระบบสำเร็จ',
      });
      const timer = setTimeout(() => {
        history.pushState('', '', '/' + cookies.get('PositionData'));
        window.location.reload(false);
      }, 1350);
    }
  };

  return (
    <Page theme={pageTheme.service}>
      <Header title={`${profile.givenName || 'to Backstage'}`}></Header>
      <Content>
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
                    เข้าสู่ระบบ
                  </Typography>
                  <form className={classes.form} noValidate>
                    <TextField
                      error={emailError ? true : false}
                      variant="outlined"
                      margin="normal"
                      required
                      fullWidth
                      id="email"
                      label="Email Address"
                      name="email"
                      value={emails}
                      onChange={emailHandleChange}
                      autoComplete="email"
                      autoFocus
                      helperText={emailError}
                    />

                    <TextField
                      error={passwordError ? true : false}
                      variant="outlined"
                      margin="normal"
                      required
                      fullWidth
                      id="password"
                      name="password"
                      label="Password"
                      type="password"
                      value={password}
                      onChange={passwordHandleChange}
                      autoComplete="current-password"
                      helperText={passwordError}
                    />

                    <Button
                      fullWidth
                      variant="contained"
                      color="primary"
                      className={classes.submit}
                      onClick={() => {
                        SigninCheck();
                      }}
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

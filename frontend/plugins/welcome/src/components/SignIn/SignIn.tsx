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
import { Cookies } from 'react-cookie/cjs'; //cookie

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
  const cookies = new Cookies();
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

  useEffect(() => {
    getPharmacist();
    getDoctor();
  }, []);

  const SigninCheck = async () => {
    var checkPharmacist = false;
    var showAlertFalse = true;

    pharmacists.map((item: any) => {
      console.log(item.email);
      if (item.email == emails && item.password == password) {
        setAlert(true);
        checkPharmacist = true;
        showAlertFalse = false;
        cookies.set('ID', JSON.stringify(item.id), { path: '/' });
        cookies.set('Name', JSON.stringify(item.name), { path: '/' });
        cookies.set('StatusLogin', JSON.stringify('Yes'), { path: '/' });
        history.pushState('', '', '/' + cookies.get('PositionData'));
        window.location.reload(false);
      }
    });

    if (checkPharmacist == false){
      doctors.map((item: any) => {
        console.log(item.email);
        if (item.email == emails && item.password == password) {
          setAlert(true);
          showAlertFalse = false;
          cookies.set('ID', JSON.stringify(item.id), { path: '/' });
          cookies.set('Name', JSON.stringify(item.name), { path: '/' });
          cookies.set('StatusLogin', JSON.stringify('Yes'), { path: '/' });
          history.pushState('', '', '/' + cookies.get('PositionData'));
          window.location.reload(false);
        }
      });
    }

    if (showAlertFalse === true){
      setAlert(false);
    }
    setStatus(true);
    const timer = setTimeout(() => {
      setStatus(false);
    }, 3000);
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

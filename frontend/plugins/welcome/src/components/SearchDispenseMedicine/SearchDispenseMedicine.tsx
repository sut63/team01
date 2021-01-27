import React, { FC, useEffect, useState } from 'react';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';

import {
  FormControl,
  TextField,
  Typography,
  Grid,
  Card,
  CardContent,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Avatar,
  Button,
} from '@material-ui/core';

import SearchOutlinedIcon from '@material-ui/icons/SearchOutlined';
import Swal from 'sweetalert2'; // alert
import moment from 'moment';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntDispenseMedicine } from '../../api/models';

import { Cookies } from 'react-cookie/cjs'; //cookie

// css style
const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    flexRow: {
      display: 'flex',
      flexWrap: 'wrap',
      flexDirection: 'row',
    },

    headLabel: {
      margin: theme.spacing(2),
    },

    formControl: {
      marginBottom: theme.spacing(2),
      maxWidth: 500,
      minWidth: 260,
    },

    tableMargin: {
      marginTop: theme.spacing(1),
    },

    button: {
      marginLeft: theme.spacing(2),
      marginBottom: theme.spacing(1),
      width: 110,
      height: 55,
    },
  }),
);

const SearchDispenseMedicine: FC<{}> = () => {
  const classes = useStyles();
  const profile = { givenName: 'ค้นหาประวัติการจ่ายยา' };
  const cookies = new Cookies();
  const api = new DefaultApi();

  //structure receive data from api
  const [apiDispenseMedicine, setApiDispenseMedicine] = useState<
    EntDispenseMedicine[]
  >([]);
  const [searchDispenseMedicine, setSearchDispenseMedicine] = useState<
    EntDispenseMedicine[]
  >([]);

  //structure show and set data
  const [sCardNumberID, setCardNumberID] = useState('');
  const [sPharmacistName, setPharmacistName] = useState(String);
  const [statusTable, setStatusTable] = useState(false);

  //structure check error
  const [validateCardNumberError, setValodateCardNumberError] = useState('');

  //function get from api
  const getDispenseMedicine = async () => {
    const res = await api.listDispensemedicine({ limit: 0, offset: 0 });
    setApiDispenseMedicine(res);
  };

  const loadingUserData = async () => {
    setPharmacistName(cookies.get('Name'));
  };

  const checkPosition = async () => {
    if (cookies.get('PositionData') != 'DispenseMedicine') {
      history.pushState('', '', '/' + cookies.get('PositionData'));
      window.location.reload(false);
    } else {
      loadingUserData();
    }
  };

  // Lifecycle Hooks
  useEffect(() => {
    checkPosition();
    getDispenseMedicine();
  }, []);

  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 2500,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

  // handleChange
  const handleChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value.toString();
    checkPattern(validateValue);
    setCardNumberID(value);
  };

  // clear input form
  function clear() {
    setCardNumberID('');
    setValodateCardNumberError('');
  }

  // ฟังก์ชั่นสำหรับ validate
  const validateText = (val: string) => {
    const regexpNum = new RegExp('^[+ 0-9]{13}$');
    return regexpNum.test(val);
  };

  // สำหรับตรวจสอบรูปแบบข้อมูลที่กรอก ว่าเป็นไปตามที่กำหนดหรือไม่
  const checkPattern = (value: string) => {
    if (validateText(value)) {
      setValodateCardNumberError('');
    } else {
      setValodateCardNumberError(
        'กรุณาใส่เลขประจําตัวประชาชนให้ครบ 13 หลัก เฉพาะตัวเลขเท่านั้น',
      );
    }
  };

  // function SearchDispenseMedicines data
  const SearchDispenseMedicines = async () => {
    if (sCardNumberID != '' && validateCardNumberError === '') {
      try {
        let res = await api.getDispensemedicine({ id: sCardNumberID });
        setSearchDispenseMedicine(res);
        setStatusTable(true);
        clear();
        Toast.fire({
          icon: 'success',
          title: 'ค้นหาข้อมูลสำเร็จ',
        });
      } catch (e) {
        Toast.fire({
          icon: 'error',
          text: 'ไม่พบข้อมูลประวัติการจ่ายยาของผู้ป่วยที่ค้นหา',
        });
      }
    } else if (validateCardNumberError != '') {
      Toast.fire({
        icon: 'error',
        text: validateCardNumberError,
      });
      setStatusTable(false);
    } else if (sCardNumberID === '') {
        Toast.fire({
          icon: 'info',
          text: 'ข้อมูลประวัติการจ่ายยาทั้งหมด',
        });
        setStatusTable(false);
    }else {
      setStatusTable(false);
    }
  };

  function switchTable() {
    if (statusTable === false) {
      return apiDispenseMedicine;
    } else {
      return searchDispenseMedicine;
    }
  }

  return (
    <Page theme={pageTheme.website}>
      <Header title={`${profile.givenName}`}>
        <Avatar src="/broken-image.jpg" />
        <div style={{ marginLeft: 10, color: 'white' }}>
          <Typography variant="h6">{sPharmacistName}</Typography>
        </div>
      </Header>

      <Content>
        <Grid
          container
          direction="column"
          justify="flex-start"
          alignItems="stretch"
        >
          <Grid item xs>
            <Typography variant="h5" style={{ marginBottom: 25 }}>
              ตารางข้อมูลประวัติการจ่ายยา
            </Typography>
            <Card>
              <CardContent>
                <Grid className={classes.flexRow}>
                  <Typography className={classes.headLabel}>
                    เลขประจําตัวประชาชน
                  </Typography>
                  <FormControl
                    variant="outlined"
                    className={classes.formControl}
                    fullWidth
                  >
                    <TextField
                      error={validateCardNumberError ? true : false}
                      name="cardNumberID"
                      label="ป้อนเลขประจําตัวประชาชน 13 หลัก"
                      helperText={validateCardNumberError}
                      value={sCardNumberID || ''}
                      onChange={handleChange}
                      variant="outlined"
                      inputProps={{ maxLength: 13 }}
                    />
                  </FormControl>
                  <Button
                    variant="contained"
                    color="secondary"
                    className={classes.button}
                    startIcon={<SearchOutlinedIcon fontSize="large" />}
                    onClick={() => {
                      SearchDispenseMedicines();
                    }}
                  >
                    <Typography variant="button">ค้นหา</Typography>
                  </Button>
                </Grid>

                <Grid className={classes.flexRow}>
                  <TableContainer
                    component={Paper}
                    className={classes.tableMargin}
                  >
                    <Table aria-label="simple table">
                      <TableHead>
                        <TableRow>
                          <TableCell align="center">วันที่</TableCell>
                          <TableCell align="center">เลขบัตรประชาชน</TableCell>
                          <TableCell align="center">ชื่อคนไข้</TableCell>
                          <TableCell align="center">แพทย์ผู้ตรวจ</TableCell>
                          <TableCell align="center">หมายเลขยา</TableCell>
                          <TableCell align="center">ยา</TableCell>
                          <TableCell align="center">แบนด์</TableCell>
                          <TableCell align="center">จำนวน</TableCell>
                          <TableCell align="center">ผู้บันทึก</TableCell>
                          <TableCell align="center">หมายเหตุ</TableCell>
                          <TableCell align="center">
                            รายละเอียดหมายเหตุ
                          </TableCell>
                          <TableCell align="center">
                            จำนวนยาที่เปลี่ยน
                          </TableCell>
                          <TableCell align="center">
                            รายละอียดยาที่เปลี่ยน
                          </TableCell>
                        </TableRow>
                      </TableHead>
                      <TableBody>
                        {switchTable().map(item => (
                          <TableRow key={item.id}>
                            <TableCell align="center">
                              {moment(item.datetime).format(
                                'DD/MM/YYYY HH.mm น.',
                              )}
                            </TableCell>
                            <TableCell align="center">
                              {
                                item.edges?.prescription?.edges
                                  ?.prescriptionpatient?.cardNumber
                              }
                            </TableCell>
                            <TableCell align="center">
                              {
                                item.edges?.prescription?.edges
                                  ?.prescriptionpatient?.name
                              }
                            </TableCell>
                            <TableCell align="center">
                              {
                                item.edges?.prescription?.edges
                                  ?.prescriptiondoctor?.name
                              }
                            </TableCell>
                            <TableCell align="center">
                              {
                                item.edges?.prescription?.edges
                                  ?.prescriptionmedicine?.serial
                              }
                            </TableCell>
                            <TableCell align="center">
                              {
                                item.edges?.prescription?.edges
                                  ?.prescriptionmedicine?.name
                              }
                            </TableCell>
                            <TableCell align="center">
                              {
                                item.edges?.prescription?.edges
                                  ?.prescriptionmedicine?.brand
                              }
                            </TableCell>
                            <TableCell align="center">
                              {item.edges?.prescription?.value}
                            </TableCell>
                            <TableCell align="center">
                              {item.edges?.pharmacist?.name}
                            </TableCell>
                            <TableCell align="center">
                              {item.edges?.annotation?.messages}
                            </TableCell>
                            <TableCell align="center">{item.note}</TableCell>
                            <TableCell align="center">
                              {item.amountchangemedicine || 0}
                            </TableCell>
                            <TableCell align="center">
                              {item.detailchangemedicine}
                            </TableCell>
                          </TableRow>
                        ))}
                      </TableBody>
                    </Table>
                  </TableContainer>
                </Grid>
              </CardContent>
            </Card>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
};
export default SearchDispenseMedicine;

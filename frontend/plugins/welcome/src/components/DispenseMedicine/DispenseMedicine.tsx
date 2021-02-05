import React, { FC, useEffect, useState } from 'react';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';

import {
  FormControl,
  MenuItem,
  TextField,
  Button,
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
  TablePagination,
  Paper,
  Avatar,
  Select,
} from '@material-ui/core';

import SaveIcon from '@material-ui/icons/Save';
import DeleteIcon from '@material-ui/icons/Delete';
import Swal from 'sweetalert2'; // alert

import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import {
  EntAnnotation,
  EntDispenseMedicine,
  EntPrescription,
} from '../../api/models';

import { Cookies } from 'react-cookie/cjs'; //cookie

// css style
const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    flexRow: {
      display: 'flex',
      flexWrap: 'wrap',
      flexDirection: 'row',
      justifyContent: 'center',
    },

    flexRowNoCen: {
      display: 'flex',
      flexWrap: 'wrap',
      flexDirection: 'row',
      justifyContent: 'space-between',
    },

    headLabel: {
      margin: theme.spacing(1),
    },

    formControl: {
      margin: 'auto',
      maxWidth: 350,
      minWidth: 260,
    },

    cardMargin: {
      marginBottom: theme.spacing(2),
    },

    tableMargin: {
      margin: theme.spacing(1),
    },

    button: {
      marginTop: theme.spacing(3),
      marginLeft: theme.spacing(1),
      marginRight: theme.spacing(1),
      width: 110,
      height: 50,
    },
  }),
);

//structure data change
interface DispenseMedicine {
  amountchangemedicine: number;
  annotation: number;
  datetime: string;
  detailchangemedicine: string;
  note: string;
  prescription: number;
}

const DispenseMedicine: FC<{}> = () => {
  const classes = useStyles();
  const profile = { givenName: 'บันทึกการจ่ายยา' };
  const cookies = new Cookies();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);

  //structure receive data from api
  const [sDispensemedicine, setDispenseMedicine] = React.useState<
    Partial<DispenseMedicine>
  >({});
  const [apiannotations, setApiAnnotation] = useState<EntAnnotation[]>([]);
  const [apiprescriptions, setApiPrescription] = useState<EntPrescription[]>(
    [],
  );
  const [apidispensemedicine, setApiDispenseMedicine] = useState<
    EntDispenseMedicine[]
  >([]);

  //structure show data
  const [sPharmacistID, setPharmacistID] = useState(Number);
  const [sPharmacistName, setPharmacistName] = useState(String);
  const [sNamePatient, setNamePatient] = useState(String);
  const [sNameDoctor, setNameDoctor] = useState(String);

  //structure table page
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(7);

  const handleChangePage = (event: unknown, newPage: number) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (
    event: React.ChangeEvent<HTMLInputElement>,
  ) => {
    setRowsPerPage(+event.target.value);
    setPage(0);
  };

  //structure check error
  const [
    amountchangemedicineInputError,
    setamountchangemedicineInputError,
  ] = React.useState(false);
  const [
    detailchangemedicineInputError,
    setdetailchangemedicineInputError,
  ] = React.useState(false);
  const [noteInputError, setnoteInputError] = React.useState(false);

  const [
    amountchangemedicineError,
    setamountchangemedicineError,
  ] = React.useState('');
  const [
    detailchangemedicineError,
    setdetailchangemedicineError,
  ] = React.useState('');
  const [noteError, setnoteError] = React.useState('');

  //function get from api
  const getAnnotation = async () => {
    const res = await api.listAnnotation({ limit: 0, offset: 0 });
    setLoading(false);
    setApiAnnotation(res);
  };

  const getDispenseMedicine = async () => {
    const res = await api.listDispensemedicine({ limit: 0, offset: 0 });
    setLoading(false);
    setApiDispenseMedicine(res);
  };

  const getPrescription = async () => {
    const res = await api.listPrescription({ limit: 0, offset: 0 });
    setLoading(false);
    setApiPrescription(res);
  };

  const loadingUserData = async () => {
    setPharmacistID(Number(cookies.get('ID')));
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
    getAnnotation();
    getDispenseMedicine();
    getPrescription();
    checkPosition();
  }, [loading]);

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
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,
  ) => {
    const name = event.target.name as keyof typeof DispenseMedicine;
    const { value } = event.target;
    const validateValue = value.toString();
    checkPattern(name, validateValue);
    setDispenseMedicine({ ...sDispensemedicine, [name]: value });
  };

  //map apiprescriptions for filter data have in apidispensemedicine
  const prescriptionMap = apiprescriptions.filter(
    presc =>
      !apidispensemedicine.some(
        dispe => dispe.edges?.prescription?.id === presc.id,
      ),
  );

  //function SetData selectPrescriptions
  function selectPrescriptions(id: any, namePatient: any, nameDoctor: any) {
    setDispenseMedicine({ ...sDispensemedicine, prescription: Number(id) });
    setNamePatient(String(namePatient));
    setNameDoctor(String(nameDoctor));
  }

  // clear input form
  function clear() {
    setDispenseMedicine({});
    setNamePatient('');
    setNameDoctor('');
  }

  // ฟังก์ชั่นสำหรับ validate ข้อความ
  const validateText = (val: string) => {
    const regexpText = new RegExp('[`~!@#$%^&*_;?<>]');
    if (regexpText.test(val)) {
      return false;
    } else {
      return true;
    }
  };

  // ฟังก์ชั่นสำหรับ validate จำนวน
  const validateAmount = (val: string) => {
    const regexpNum = new RegExp('^[+ 0-9]{0,3}$');
    return regexpNum.test(val);
  };

  // สำหรับตรวจสอบรูปแบบข้อมูลที่กรอก ว่าเป็นไปตามที่กำหนดหรือไม่
  const checkPattern = (name: string, value: string) => {
    switch (name) {
      case 'note':
        validateText(value)
          ? setnoteError('')
          : setnoteError(
              'ห้ามป้อนอักษรพิเศษ ถ้าไม่มีรายละเอียดหมายเหตุ กรุณาใส่เครื่องหมาย " - "',
            );
        validateText(value)
          ? setnoteInputError(false)
          : setnoteInputError(true);
        return;
      case 'amountchangemedicine':
        validateAmount(value)
          ? setamountchangemedicineError('')
          : setamountchangemedicineError(
              'ถ้าไม่มีจำนวนยาที่ถูกเปลี่ยน กรุณาใส่เลข " 0 "',
            );
        validateAmount(value)
          ? setamountchangemedicineInputError(false)
          : setamountchangemedicineInputError(true);
        return;
      case 'detailchangemedicine':
        validateText(value)
          ? setdetailchangemedicineError('')
          : setdetailchangemedicineError(
              'ห้ามป้อนอักษรพิเศษ ถ้าไม่มีรายละเอียดยาที่ถูกเปลี่ยน กรุณาใส่เครื่องหมาย " - "',
            );
        validateText(value)
          ? setdetailchangemedicineInputError(false)
          : setdetailchangemedicineInputError(true);
        return;
      default:
        return;
    }
  };

  const alertMessage = (icon: any, text: any) => {
    Toast.fire({
      icon: icon,
      title: 'บันทึกข้อมูลไม่สำเร็จ',
      text: text,
    });
  };

  const checkCaseSaveError = (err: string) => {
    switch (err) {
      case 'pharmacist not found':
        alertMessage('error', 'ไม่พบข้อมูลเภสัชกร');
        return;

      case 'prescription not found':
        alertMessage('error', 'กรุณาเลือกคิวที่รับการจ่ายยา');
        return;

      case 'datetime not found':
        alertMessage('error', 'กรุณาเลือกวันที่');
        return;

      case 'annotation not found':
        alertMessage('error', 'กรุณาเลือกหมายเหตุ');
        return;

      case 'note':
        alertMessage(
          'error',
          'ถ้าไม่มีรายละเอียดหมายเหตุ กรุณาใส่เครื่องหมาย " - "',
        );
        return;

      case 'amountchangemedicine':
        alertMessage('error', 'ถ้าไม่มีจำนวนยาที่ถูกเปลี่ยน กรุณาใส่เลข "0"');
        return;

      case 'detailchangemedicine':
        alertMessage(
          'error',
          'ถ้าไม่มีรายละเอียดยาที่ถูกเปลี่ยน กรุณาใส่เครื่องหมาย " - "',
        );
        return;
      default:
        return;
    }
  };

  // function CreateDispenseMedicines data
  const CreateDispenseMedicines = async () => {
    const dispensemediciness = {
      amountchangemedicine: Number(sDispensemedicine.amountchangemedicine),
      annotation: sDispensemedicine.annotation,
      datetime: sDispensemedicine.datetime + ':00+07:00',
      detailchangemedicine: sDispensemedicine.detailchangemedicine,
      note: sDispensemedicine.note,
      pharmacist: sPharmacistID,
      prescription: sDispensemedicine.prescription,
    };

    const api = 'http://localhost:8080/api/v1/dispensemedicines';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(dispensemediciness),
    };

    fetch(api, requestOptions)
      .then(response => response.json())
      .then(data => {
        if (data.status === true) {
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
          setLoading(true);
          clear();
        } else {
          if (data.error.Name === undefined) {
            checkCaseSaveError(data.error);
          } else {
            checkCaseSaveError(data.error.Name);
          }
        }
      });
  };

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
          direction="row"
          justify="space-evenly"
          alignItems="flex-start"
        >
          <Grid item xs>
            <Typography variant="h5" style={{ marginBottom: 25 }}>
              ตารางข้อมูลคิว
            </Typography>
            <Paper>
              <TablePagination
                rowsPerPageOptions={[5, 7, 15]}
                component="div"
                count={prescriptionMap.length}
                rowsPerPage={rowsPerPage}
                page={page}
                onChangePage={handleChangePage}
                onChangeRowsPerPage={handleChangeRowsPerPage}
              />
              <TableContainer component={Paper}>
                <Table aria-label="simple table">
                  <TableHead>
                    <TableRow>
                      <TableCell align="center">No.</TableCell>
                      <TableCell align="center">รายชื่อคนไข้</TableCell>
                      <TableCell align="center">เลือกคิว</TableCell>
                    </TableRow>
                  </TableHead>
                  <TableBody>
                    {prescriptionMap
                      .slice(
                        page * rowsPerPage,
                        page * rowsPerPage + rowsPerPage,
                      )
                      .map((item: any) => (
                        <TableRow key={item.id}>
                          <TableCell align="center">{item.id}</TableCell>
                          <TableCell align="center">
                            {item.edges?.prescriptionpatient?.name}
                          </TableCell>
                          <TableCell align="center">
                            <Button
                              onClick={() => {
                                selectPrescriptions(
                                  item.id,
                                  item.edges?.prescriptionpatient?.name,
                                  item.edges?.prescriptiondoctor?.name,
                                );
                              }}
                              style={{ marginLeft: 2 }}
                              variant="contained"
                              color="primary"
                            >
                              เลือก
                            </Button>
                          </TableCell>
                        </TableRow>
                      ))}
                  </TableBody>
                </Table>
              </TableContainer>
            </Paper>
          </Grid>

          <Grid
            item
            xs={7}
            style={{
              minWidth: 400,
            }}
          >
            <Typography variant="h5" style={{ marginBottom: 25 }}>
              ข้อมูลใบสั่งยา
            </Typography>
            <Card className={classes.cardMargin}>
              <CardContent>
                <Grid className={classes.flexRow}>
                  <FormControl
                    variant="outlined"
                    className={classes.formControl}
                    style={{
                      maxWidth: 200,
                      minWidth: 100,
                    }}
                  >
                    <Typography className={classes.headLabel}>
                      ลำดับที่
                    </Typography>
                    <TextField
                      label="เลขคิว"
                      name="queue"
                      variant="outlined"
                      value={sDispensemedicine.prescription || ''} // (undefined || '') = ''
                      inputProps={{ readOnly: true }}
                    ></TextField>
                  </FormControl>

                  <FormControl
                    variant="outlined"
                    className={classes.formControl}
                    fullWidth
                  >
                    <Typography className={classes.headLabel}>คนไข้</Typography>
                    <TextField
                      label="ชื่อ-นามสกุล"
                      name="name"
                      variant="outlined"
                      value={sNamePatient || ''} // (undefined || '') = ''
                      inputProps={{ readOnly: true }}
                    ></TextField>
                  </FormControl>

                  <FormControl
                    variant="outlined"
                    className={classes.formControl}
                    fullWidth
                  >
                    <Typography className={classes.headLabel}>
                      แพทย์ผู้ตรวจ
                    </Typography>
                    <TextField
                      label="ชื่อ-นามสกุล"
                      name="physician"
                      variant="outlined"
                      value={sNameDoctor || ''} // (undefined || '') = ''
                      inputProps={{ readOnly: true }}
                    ></TextField>
                  </FormControl>
                </Grid>
              </CardContent>
            </Card>

            <Card className={classes.cardMargin}>
              <CardContent>
                <Grid className={classes.flexRow}>
                  <TableContainer
                    component={Paper}
                    className={classes.tableMargin}
                  >
                    <Table aria-label="simple table">
                      <TableHead>
                        <TableRow>
                          <TableCell align="center">No.</TableCell>
                          <TableCell align="center">ชื่อยา</TableCell>
                          <TableCell align="center">หมายเลขยา</TableCell>
                          <TableCell align="center">แบนด์</TableCell>
                          <TableCell align="center">จำนวน</TableCell>
                          <TableCell align="center">
                            จำนวนที่เหลือในคลัง
                          </TableCell>
                        </TableRow>
                      </TableHead>
                      <TableBody>
                        {apiprescriptions
                          .filter(
                            item => item.id === sDispensemedicine.prescription,
                          )
                          .map((item: any) => (
                            <TableRow key={item.id}>
                              <TableCell align="center">
                                {item.edges?.prescriptionmedicine?.id}
                              </TableCell>
                              <TableCell align="center">
                                {item.edges?.prescriptionmedicine?.name}
                              </TableCell>
                              <TableCell align="center">
                                {item.edges?.prescriptionmedicine?.serial}
                              </TableCell>
                              <TableCell align="center">
                                {item.edges?.prescriptionmedicine?.brand}
                              </TableCell>
                              <TableCell align="center">{item.value}</TableCell>
                              <TableCell align="center">
                                {item.edges?.prescriptionmedicine?.amount -
                                  item.value}
                              </TableCell>
                            </TableRow>
                          ))}
                      </TableBody>
                    </Table>
                  </TableContainer>
                </Grid>
              </CardContent>
            </Card>

            <Card className={classes.cardMargin}>
              <CardContent>
                <Grid className={classes.flexRow}>
                  <FormControl
                    variant="outlined"
                    className={classes.formControl}
                  >
                    <Typography className={classes.headLabel}>
                      วันที่
                    </Typography>
                    <TextField
                      label="วันที่"
                      name="datetime"
                      type="datetime-local"
                      value={sDispensemedicine.datetime || ''} // (undefined || '') = ''
                      onChange={handleChange}
                      InputLabelProps={{
                        shrink: true,
                      }}
                    />
                  </FormControl>

                  <FormControl
                    variant="outlined"
                    className={classes.formControl}
                    fullWidth
                  >
                    <Typography className={classes.headLabel}>
                      หมายเหตุ
                    </Typography>

                    <Select
                      labelId="annotation"
                      label="เลือกหมายเหตุ"
                      name="annotation"
                      value={sDispensemedicine.annotation || ''} // (undefined || '') = ''
                      onChange={handleChange}
                    >
                      {apiannotations.map(option => (
                        <MenuItem key={option.id} value={option.id}>
                          {option.messages}
                        </MenuItem>
                      ))}
                    </Select>
                  </FormControl>

                  <FormControl
                    variant="outlined"
                    className={classes.formControl}
                  >
                    <Typography className={classes.headLabel}>
                      รายละเอียดหมายเหตุ
                    </Typography>
                    <TextField
                      error={noteInputError}
                      name="note"
                      label="ป้อนรายละเอียดหมายเหตุ"
                      helperText={noteError}
                      value={sDispensemedicine.note || ''}
                      onChange={handleChange}
                      variant="outlined"
                    />
                  </FormControl>

                  <FormControl
                    variant="outlined"
                    className={classes.formControl}
                  >
                    <Typography className={classes.headLabel}>
                      จำนวนยาที่ถูกเปลี่ยน
                    </Typography>
                    <TextField
                      error={amountchangemedicineInputError}
                      name="amountchangemedicine"
                      label="ป้อนจำนวนยาที่ถูกเปลี่ยน"
                      type="number"
                      helperText={amountchangemedicineError}
                      value={sDispensemedicine.amountchangemedicine || ''}
                      onChange={handleChange}
                      variant="outlined"
                    />
                  </FormControl>

                  <FormControl
                    variant="outlined"
                    className={classes.formControl}
                  >
                    <Typography className={classes.headLabel}>
                      รายละเอียดยาที่ถูกเปลี่ยน
                    </Typography>
                    <TextField
                      error={detailchangemedicineInputError}
                      name="detailchangemedicine"
                      label="ป้อนรายละเอียดยาที่ถูกเปลี่ยน"
                      helperText={detailchangemedicineError}
                      value={sDispensemedicine.detailchangemedicine || ''}
                      onChange={handleChange}
                      variant="outlined"
                    />
                  </FormControl>

                  <FormControl
                    variant="outlined"
                    className={classes.flexRow}
                    fullWidth
                  >
                    <Button
                      variant="contained"
                      color="secondary"
                      className={classes.button}
                      startIcon={<SaveIcon />}
                      onClick={() => {
                        CreateDispenseMedicines();
                      }}
                    >
                      <Typography variant="button">บันทึก</Typography>
                    </Button>

                    <Button
                      variant="contained"
                      className={classes.button}
                      startIcon={<DeleteIcon />}
                      onClick={() => {
                        clear();
                      }}
                    >
                      <Typography variant="button">ล้างค่า</Typography>
                    </Button>
                  </FormControl>
                </Grid>
              </CardContent>
            </Card>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
};
export default DispenseMedicine;

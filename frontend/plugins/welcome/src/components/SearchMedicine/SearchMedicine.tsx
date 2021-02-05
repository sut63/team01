import React, { FC, useState, useEffect, useCallback } from 'react';
import { makeStyles, ThemeProvider, Theme } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import FormControl from '@material-ui/core/FormControl';
import { DefaultApi } from '../../api/apis';
import { Cookies } from 'react-cookie/cjs'; //cookie
import SearchOutlinedIcon from '@material-ui/icons/SearchOutlined';
import { EntMedicine } from '../../api/models/EntMedicine';

import {
  Grid,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Avatar,
  Button,
  Card,
  CardContent,
  IconButton,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
} from '@material-ui/core';

// header css
const HeaderCustom = {
  minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
  table: {
    flexGrow: 1,
  },
  formControl: {
    margin: theme.spacing(1),
    minWidth: 520,
  },
  card: {
    width: 1300,
  },
  margin: {
    margin: theme.spacing(3),
  }
}));

const SearchMedicine: FC<{}> = () => {
  const classes = useStyles();
  const api = new DefaultApi();
  const cookies = new Cookies();
  const [PhaName, setPhaName] = React.useState("");
  const [Serial, setSerial] = React.useState("");
  const [Medicine, setMedicine] = React.useState<EntMedicine[]>([]);
  const [SearchMedicine, setSearchMedicine] = React.useState<EntMedicine[]>([]);
  const [statusTable, setStatusTable] = React.useState(false);

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

  const handleChange = (event: React.ChangeEvent<{ name?: string; value: any }>) => {
    const { value } = event.target;
    setSerial(value);
  };

  const getMedicines = async () => {
    const res = await api.listMedicine({ limit: 20, offset: 0 });
    setMedicine(res);
  };

  const search = async () => {
    if (Serial != "") {
      try {
        let res = await api.getMedicine({ id: Serial });
        setSearchMedicine(res);
        setStatusTable(true);
        clear();
        Toast.fire({
          icon: 'success',
          title: 'ค้นหาข้อมูลสำเร็จ',
        });
      } catch (e) {
        Toast.fire({
          icon: 'error',
          title: "ไม่พบข้อมูลของยาที่ทำการค้นหา"
        });
        setStatusTable(false);
      }
    }
    else {
      Toast.fire({
        icon: 'info',
        title: "แสดงข้อมูลทั้งหมด"
      });
      setStatusTable(false);
    }
  }

  const clear = () => {
    setSerial("");
  }

  function checkTable() {
    if (statusTable === false) {
      return Medicine
    }
    else {
      return SearchMedicine
    }
  }

  const checkPosition = async () => {
    if (cookies.get('PositionData') != 'Medicine') {
      history.pushState('', '', '/' + cookies.get('PositionData'));
      window.location.reload(false);
    }
  };

  useEffect(() => {
    getMedicines();
    checkPosition();
    setPhaName(cookies.get('Name'));
  }, []);

  return (<Page theme={pageTheme.service}>
    <Header style={HeaderCustom} title={'ค้นหาข้อมูลยา'}>
      <Avatar src="/broken-image.jpg" />
      <div style={{ marginLeft: 10 }}>{PhaName}</div>
    </Header>
    <Content>
      <Grid
        container
        direction="column"
        justify="center"
        alignItems="center"

      >
        <div>
          <Card className={classes.card} style={{ marginTop: 10 }}>
            <CardContent>
              <div>
                <TextField
                  style={{ width: 500, marginLeft: 25 }}
                  name="Serial"
                  label="กรุณาระบุรหัสยา"
                  onChange={handleChange}
                  variant="outlined"
                  value={Serial || ''}
                  inputProps={{ maxLength: 13 }}
                />
                <IconButton aria-label="search" style={{ marginLeft: 20 }} onClick={search}>
                  <SearchOutlinedIcon fontSize="large" />
                </IconButton>
              </div>
              <div>
                <TableContainer component={Paper}>
                  <Table className={classes.table} aria-label="simple table">
                    <TableHead>
                      <TableRow>
                        <TableCell align="center">ID.</TableCell>
                        <TableCell align="center">ชื่อยา</TableCell>
                        <TableCell align="center">รหัสยา</TableCell>
                        <TableCell align="center">ยี่ห้อยา</TableCell>
                        <TableCell align="center">จำนวนยา</TableCell>
                        <TableCell align="center">ราคายา</TableCell>
                        <TableCell align="center">วิธีการใช้ยา</TableCell>
                        <TableCell align="center">ระดับความอันตราย</TableCell>
                        <TableCell align="center">ประเภทยา</TableCell>
                        <TableCell align="center">หน่วยยา</TableCell>
                      </TableRow>
                    </TableHead>
                    <TableBody>
                      {checkTable().map(item => (
                        <TableRow key={item.id}>
                          <TableCell align="center">{item.id}</TableCell>
                          <TableCell align="center">{item.name}</TableCell>
                          <TableCell align="center">{item.serial}</TableCell>
                          <TableCell align="center">{item.brand}</TableCell>
                          <TableCell align="center">{item.amount}</TableCell>
                          <TableCell align="center">{item.price}</TableCell>
                          <TableCell align="center">{item.howtouse}</TableCell>
                          <TableCell align="center">{item.edges?.levelOfDangerous?.name}</TableCell>
                          <TableCell align="center">{item.edges?.medicineType?.name}</TableCell>
                          <TableCell align="center">{item.edges?.unitOfMedicine?.name}</TableCell>
                        </TableRow>
                      ))}
                    </TableBody>
                  </Table>
                </TableContainer>
              </div>
            </CardContent>
          </Card>
        </div>
      </Grid>
    </Content>
  </Page>
  );
};

export default SearchMedicine;

import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import Typography from '@material-ui/core/Typography';
import Divider from '@material-ui/core/Divider';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow'
import Paper from '@material-ui/core/Paper';
import moment from 'moment';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import {
    MuiPickersUtilsProvider,
    KeyboardTimePicker,
    KeyboardDatePicker,
} from '@material-ui/pickers';
import {
    Container,
    Grid,
    FormControl,
    Select,
    InputLabel,
    MenuItem,
    TextField,
    Avatar,
    Button,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntMedicine, EntPharmacist, EntCompany, EntOrder } from '../../api';
import { Cookies } from 'react-cookie/cjs'; //cookie
import Snackbar from '@material-ui/core/Snackbar';
import Alert from '@material-ui/lab/Alert';
// header css
const HeaderCustom = {
    minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
    root: {
        flexGrow: 1,
    },
    paper: {
        marginTop: theme.spacing(2),
        marginBottom: theme.spacing(2),
    },
    formControl: {
        width: 300,
    },
    selectEmpty: {
        marginTop: theme.spacing(2),
    },
    container: {
        display: 'flex',
        flexWrap: 'wrap',
    },
    textField: {
        width: 300,
    },
    bottom: {
        marginLeft: theme.spacing(3),
        marginTop: theme.spacing(1),
        marginBottom: theme.spacing(2),
    },
    divider: {
        margin: theme.spacing(2, 0),
    },
    table: {
        minWidth: 650,
    },
}));


const SearchOrder: FC<{}> = () => {
    const classes = useStyles();
    const api = new DefaultApi();

    // ดึงคุกกี้
    const cookies = new Cookies();
    const [PhaName, setPhaName] = React.useState("");

    // Customer
    const [idMedicine, setIdMedicine] = React.useState<number>(0)
    const [medicine, setMedicine] = React.useState<EntMedicine[]>([])
    const getMedicine = async () => {
        const res = await api.listMedicine({})
        setMedicine(res)
    }

    // alert setting
    const [open, setOpen] = React.useState(false);
    const [fail, setFail] = React.useState(false);

    //close alert 
    const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
        if (reason === 'clickaway') {
            return;
        }
        setFail(false);
        setOpen(false);
    };


    
    var lenorder : number
    const [orders, setOrders] = React.useState<EntOrder[]>([])
     const getOrders = async () => {
        const res = await api.getOrder({ id:idMedicine})
        setOrders(res)
        lenorder = res.length
        if (lenorder > 0){
            setOpen(true)
        }else{
            setFail(true)
        }   
    }
    const listOrders = async () => {
        const res = await api.listOrder({})
        setOrders(res)
    }

    // Lifecycle Hooks
    useEffect(() => {
        getMedicine();
        listOrders();
    }, []);


    // set data to object and setIdcustomer
    const handleChange = (
        event: React.ChangeEvent<{ name?: string; value: any }>,
    ) => {
        const { value } = event.target;
        setIdMedicine(value);
    };

    // clear cookies
    function Clears() {
        setPhaName(cookies.get('Name'));
        window.location.reload(false)
    }
    console.log(orders)
    // function seach data
   function seach() {
        getOrders();
    }

    return (
        <Page theme={pageTheme.home}>
            <Header style={HeaderCustom} title={`ระบบค้นหาการสั่งซื้อยาย้อนหลัง`}>
                <Avatar alt="Remy Sharp" src="../../image/account.jpg" />
                <div style={{ marginLeft: 10, marginRight: 20 }}>{PhaName}</div>
               
            </Header>
            <Content>
                <Grid container spacing={1}>
                    <Grid item xs={1}>
                        <div className={classes.paper}><h3>Medicine</h3></div>
                    </Grid>
                    <Grid item xs={3}>
                        <FormControl variant="outlined" className={classes.formControl}>
                            <InputLabel>select medicine</InputLabel>
                            <Select
                                name="Medicine"
                                value={idMedicine || ''} // (undefined || '') = ''
                                onChange={handleChange}
                            >
                                {medicine.map(item => {
                                    return (
                                        <MenuItem key={item.id} value={item.id}>
                                            {item.name}
                                        </MenuItem>
                                    );
                                })}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                        <Button
                            variant="contained"
                            color="primary"
                            size="large"
                            onClick={seach}
                        >
                            seach
                        </Button>
                    </Grid>
                </Grid>

                <Grid item xs={12}>
                    <Divider className={classes.divider} />
                    <Typography variant="subtitle1" gutterBottom>
                        Order Tables:
                        </Typography>
                </Grid>
                <TableContainer component={Paper}>
                    <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="center">id</TableCell>
                                <TableCell align="center">Pharmacist</TableCell>
                                <TableCell align="center">Medicine</TableCell>
                                <TableCell align="center">amount</TableCell>
                                <TableCell align="center">price</TableCell>
                                <TableCell align="center">phone</TableCell>
                                <TableCell align="center">company</TableCell>
                                <TableCell align="center">date</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {orders.map((item: EntOrder) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.id}</TableCell>
                                    <TableCell align="center">{item.edges?.pharmacist?.name}</TableCell>
                                    <TableCell align="center">{item.edges?.medicine?.name}</TableCell>
                                    <TableCell align="center">{item.amount}</TableCell>
                                    <TableCell align="center">{item.price}</TableCell>
                                    <TableCell align="center">{item.hospitalid}</TableCell>
                                    <TableCell align="center">{item.edges?.company?.name}</TableCell>
                                    <TableCell align="center">{moment(item.addedtime).format('DD/MM/YYYY HH:mm:ss')}</TableCell>
                                </TableRow> 
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
                <Snackbar open={open} autoHideDuration={3000} onClose={handleClose}>
            <Alert onClose={handleClose} >
                
                
              ค้นหาสำเร็จ
          </Alert>
          </Snackbar>

          <Snackbar open={fail} autoHideDuration={3000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="error">
              ไม่พบข้อมูล
          </Alert>
          </Snackbar>
            </Content>
        </Page>
    );
};

export default SearchOrder;

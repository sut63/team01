import React, { FC, useEffect, useCallback } from 'react';
import { makeStyles, ThemeProvider, Theme } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import FormControl from '@material-ui/core/FormControl';
import { DefaultApi } from '../../api/apis';
import { Cookies } from 'react-cookie/cjs'; //cookie
import SearchOutlinedIcon from '@material-ui/icons/SearchOutlined';

import { EntPatientInfo } from '../../api/models/EntPatientInfo'
import { EntMedicine } from '../../api/models/EntMedicine'
import { EntDrugAllergy } from '../../api/models/EntDrugAllergy'

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
    root: {
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

const SearchDrugAllergy: FC<{}> = () => {
    const classes = useStyles();
    const api = new DefaultApi();
    const cookies = new Cookies();
    const [PhaName, setPhaName] = React.useState("");
    const [cardNumber, setCardNumber] = React.useState("");
    const [validateCardNumberText, setValodateCardNumberText] = React.useState("");
    const [drugAllergy, setDrugAllergy] = React.useState<EntDrugAllergy[]>([]);
    const [searchDrugAllergy, setSearchDrugAllergy] = React.useState<EntDrugAllergy[]>([]);
    const [statusTable, setStatusTable] = React.useState(false);
    const [validateCardNumberInput, setValidateCardNumberInput] = React.useState(false);
    //const [storage, setStorage] = useState<EntStorage[]>([]);

    // alert setting
    const Toast = Swal.mixin({
        //toast: true,
        position: 'center',
        showConfirmButton: false,
        timer: 1000,
        timerProgressBar: true,
        didOpen: toast => {
            toast.addEventListener('mouseenter', Swal.stopTimer);
            toast.addEventListener('mouseleave', Swal.resumeTimer);
        },
    });

    const handleChange = (event: React.ChangeEvent<{ name?: string; value: any }>) => {
        const { value } = event.target;
        checkValidate(value);
        setCardNumber(value);
    };

    const getDrugAllergy = async () => {
        const res = await api.listDrugAllergy({ limit: 20, offset: 0 });
        setDrugAllergy(res);
    };

    const checkValidate = (text: string) => {
        if(text.length > 12 || text === ""){
            setValodateCardNumberText("")
            setValidateCardNumberInput(false);
        }
        else{
            setValodateCardNumberText("กรุณาใส่เลขประจําตัวประชาชนให้ครบ 13 หลัก");
            setValidateCardNumberInput(true);
        }
    }

    const search = async () => {
        if ((cardNumber != "") && (validateCardNumberInput === false)) {
            try {
                let res = await api.getDrugAllergy({ card: cardNumber });
                setSearchDrugAllergy(res);
                setStatusTable(true);
                clear();
            } catch (e) {
                Toast.fire({
                    icon: 'error',
                    text: "ไม่พบข้อมูลประวัติการเเพ้ยาของผู้ป่วยที่ค้นหา"
                });
                setStatusTable(false);
            }
        }
        else if(validateCardNumberInput === true){
            Toast.fire({
                icon: 'error',
                text: validateCardNumberText,
            });
            setStatusTable(false);
        }
        else {
            setStatusTable(false);
        }
    }

    const clear = () => {
        setCardNumber("");
        setValodateCardNumberText("")
        setValidateCardNumberInput(false);
    }

    function checkTable() {
        if (statusTable === false) {
            return drugAllergy
        }
        else {
            return searchDrugAllergy
        }
    }

    const checkPosition = async () => {
      if (cookies.get('PositionData') != 'DrugAllergy') {
        history.pushState('', '', '/' + cookies.get('PositionData'));
        window.location.reload(false); 
      }
    };

    useEffect(() => {
    }, [statusTable]);

    useEffect(() => {
        getDrugAllergy();
        checkPosition();
        setPhaName(cookies.get('Name'));
    }, []);

    return (
        <Page theme={pageTheme.service}>
            <Header style={HeaderCustom} title={'ค้นหาประวัติการเเพ้'}>
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
                                        error={validateCardNumberInput}
                                        style={{ width: 500, marginLeft: 25 }}
                                        name="cardNumber"
                                        label="เลขประจําตัวประชาชน"
                                        onChange={handleChange}
                                        variant="outlined"
                                        value={cardNumber || ''}
                                        helperText={validateCardNumberText}
                                        inputProps={{ maxLength: 13 }}
                                    />
                                    <IconButton aria-label="search" style={{ marginLeft: 20 }} onClick={search}>
                                        <SearchOutlinedIcon fontSize="large" />
                                    </IconButton>
                                </div>
                                <div>
                                    <TableContainer component={Paper} style={{ marginTop: 40, marginLeft: 30, width: 1200 }}>
                                        <Table  aria-label="simple table">
                                            <TableHead>
                                                <TableRow>
                                                    <TableCell style={{ width: 150 }}>วันที่</TableCell>
                                                    <TableCell style={{ width: 120 }}>เลขบัตรประชาชน</TableCell>
                                                    <TableCell style={{ width: 150 }}>ชื่อผู้ป่วย</TableCell>
                                                    <TableCell style={{ width: 120 }}>ยาที่เเพ้</TableCell>
                                                    <TableCell style={{ width: 150 }}>อาการเเพ้</TableCell>
                                                    <TableCell style={{ width: 150 }}>ผู้บันทึก</TableCell>
                                                </TableRow>
                                            </TableHead>
                                            <TableBody>
                                                {checkTable().map(item => (
                                                    <TableRow key={item.id}>
                                                        <TableCell >{item.dateTime}</TableCell>
                                                        <TableCell >{item.edges?.patient?.cardNumber}</TableCell>
                                                        <TableCell >{item.edges?.patient?.name}</TableCell>
                                                        <TableCell >{item.edges?.medicine?.name}</TableCell>
                                                        <TableCell >{item.symptom}</TableCell>
                                                        <TableCell >{item.edges?.pharmacist?.name}</TableCell>
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

export default SearchDrugAllergy;

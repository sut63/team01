import React, { FC, useEffect } from 'react';
import { makeStyles, ThemeProvider, Theme } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import ArrowBackIcon from '@material-ui/icons/ArrowBack';
import SearchIcon from '@material-ui/icons/Search';
import FormControl from '@material-ui/core/FormControl';
import { DefaultApi } from '../../api/apis';

import { EntPatientInfo } from '../../api/models/EntPatientInfo'
import { EntMedicine } from '../../api/models/EntMedicine'

import {
    Container,
    Grid,
    Select,
    InputLabel,
    MenuItem,
    TextField,
    Avatar,
    Button,
    Card,
    CardActions,
    CardContent,
    CardHeader,
    IconButton,
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
        width: 600,
    },
    margin: {
        margin: theme.spacing(3),
    }
}));

interface DrugAllergy {
    patient: number;
    medicine: number;
    pharmacist: number;
    dateTime: Date;
}

const DrugAllergy: FC<{}> = () => {
    const classes = useStyles();
    const api = new DefaultApi();
    const [drugAllergy, setDrugAllergy] = React.useState<Partial<DrugAllergy>>({});
    const [patients, setPatients] = React.useState<EntPatientInfo[]>([]);
    const [medicine, setMedicine] = React.useState<EntMedicine[]>([]);
    const [loading, setLoading] = React.useState(true);
    

    const handleChange = (event: React.ChangeEvent<{ name?: string; value: unknown }>) => {
        const name = event.target.name as keyof typeof DrugAllergy;
        const { value } = event.target;
        setDrugAllergy({ ...drugAllergy, [name]: value });
        console.log(drugAllergy);
    };

    const getPatient = async () => {
        const res = await api.listPatientInfo({ limit: 10, offset: 0 });
        setLoading(false);
        setPatients(res);
    };

    const getMedicine = async () => {
        const res = await api.listMedicine({ limit: 10, offset: 0 });
        setLoading(false);
        setMedicine(res);
    };

    const save = async () => {
        console.log(drugAllergy);
    };


    // Lifecycle Hooks
    useEffect(() => {
        getPatient();
        getMedicine();
    }, []);

    return (
        <Page theme={pageTheme.service}>
            <Header style={HeaderCustom} title={'บันทึกประวัติการเเพ้ยา'}>
                <Avatar src="/broken-image.jpg" />
                <div style={{ marginLeft: 10 }}>Name Surname</div>
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
                                <div className={classes.margin}>
                                    <div>
                                        <TextField
                                            style={{ width: 520 }}
                                            name="pharmacist"
                                            label="เภสัชกร"
                                            defaultValue="Name Surname"
                                            InputProps={{
                                                readOnly: true,
                                            }}
                                            variant="outlined"
                                        />
                                    </div>
                                    <FormControl variant='outlined' style={{ width: 520, marginTop: 40 }}>
                                        <InputLabel id="patient">ผู้ป่วย</InputLabel>
                                        <Select
                                            labelId="patient"
                                            name="patient"
                                            onChange={handleChange}
                                            label='ผู้ป่วย'
                                        >
                                            {patients.map(item => {
                                                return (
                                                    <MenuItem key={item.id} value={item.id}>
                                                        {item.name}
                                                    </MenuItem>
                                                );
                                            })}
                                        </Select>
                                    </FormControl>
                                    <FormControl variant='outlined' style={{ width: 520, marginTop: 40 }}>
                                        <InputLabel id="medicine">ยา</InputLabel>
                                        <Select
                                            labelId="medicine"
                                            name="medicine"
                                            onChange={handleChange}
                                            label='ยา'
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
                                    <div style={{ marginTop: 40 }}>
                                        <TextField
                                            label="เลือกวันที่"
                                            name="dateTime"
                                            type="datetime-local"
                                            variant='outlined'
                                            InputLabelProps={{
                                                shrink: true,
                                            }}
                                            style={{ width: 520 }}
                                            onChange={handleChange}
                                        />
                                    </div>
                                </div>
                            </CardContent>
                        </Card>
                    </div>
                    <div style={{ marginTop: 20 }}>
                        <Button
                            variant="contained"
                            color="primary"
                            size="large"
                            startIcon={<SaveIcon />}
                            style={{ width: 150 }}
                            onClick={ save }
                        >
                            บันทึก
                        </Button>
                        <Button
                            variant="contained"
                            color="secondary"
                            size="large"
                            startIcon={<ArrowBackIcon fontSize="small" />}
                            style={{ marginLeft: 20, width: 150 }}
                        >
                            ย้อนกลับ
                        </Button>
                    </div>
                </Grid>
            </Content>
        </Page>
    );
};

export default DrugAllergy;

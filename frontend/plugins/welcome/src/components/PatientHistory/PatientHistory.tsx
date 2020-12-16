import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import ArrowBackIcon from '@material-ui/icons/ArrowBack';

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
    Card,
    CardActions,
    CardContent,
    CardHeader,
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
    card: {
        width: 900,
    },
}));

const PatientHistory: FC<{}> = () => {
    const classes = useStyles();

    const [phoneNumber, setPhoneNumber] = React.useState('')

    const handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setPhoneNumber(event.target.value as string)
    };

    // Lifecycle Hooks
    useEffect(() => {

    }, []);

    return (
        <Page theme={pageTheme.service}>
            <Header style={HeaderCustom} title={'บันทึกประวัติผู้ป่วย'}>
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
                        <Card className={classes.card}>
                            <CardHeader title='ข้อมูลทั่วไป' />
                            <CardContent>
                                <div>
                                    <TextField
                                        required
                                        label="เลขบัตรประชาชน 13 หลัก"
                                        variant="outlined"
                                        style={{ width: 415 }}
                                        inputProps={{ maxLength: 13 }}
                                        helperText="ถ้ามีข้อมูลอยู่เเล้วระบบจะกรอกข้อมูลทั่วไปให้อัตโนมัติ"
                                    />
                                    <div style={{ marginTop: 20 }}>
                                        <FormControl style={{ width: 95 }}>
                                            <InputLabel>คํานําหน้า</InputLabel>
                                            <Select
                                                id="blood_type"
                                                onChange={handleChange}
                                                variant='outlined'
                                            >
                                            </Select>
                                        </FormControl>
                                        <TextField
                                            required
                                            label="ชื่อ"
                                            variant="outlined"
                                            style={{ width: 300, marginLeft: 20 }}
                                        />
                                        <TextField
                                            required
                                            label="นามสกุล"
                                            variant="outlined"
                                            style={{ width: 300, marginLeft: 20 }}
                                        />
                                    </div>
                                    <div style={{ marginTop: 20 }}>
                                        <FormControl style={{ width: 95 }}>
                                            <InputLabel required>เพศ</InputLabel>
                                            <Select
                                                id="gender"
                                                onChange={handleChange}
                                                variant='outlined'
                                            >
                                            </Select>
                                        </FormControl>
                                        <TextField
                                            required
                                            label="อายุ"
                                            variant="outlined"
                                            style={{ width: 95, marginLeft: 20 }}
                                        />
                                        <FormControl style={{ width: 95, marginLeft: 20 }}>
                                            <InputLabel required>หมู่เลือด</InputLabel>
                                            <Select
                                                id="blood_type"
                                                onChange={handleChange}
                                                variant='outlined'
                                            >
                                            </Select>
                                        </FormControl>
                                    </div>
                                    <div style={{ marginTop: 20 }}>
                                        <TextField
                                            label="โรคประจําตัว"
                                            variant="outlined"
                                            style={{ width: 325 }}
                                            multiline
                                            rows={3}
                                        />
                                        <TextField
                                            label="โรคเรื้อรัง"
                                            variant="outlined"
                                            style={{ width: 325, marginLeft: 20 }}
                                            multiline
                                            rows={3}
                                        />
                                    </div>
                                    <div style={{ marginTop: 20 }}>
                                        <TextField
                                            value={phoneNumber}
                                            label="บ้านเเลขที่"
                                            variant="outlined"
                                            onChange={handleChange}
                                            style={{ width: 95 }}
                                        />
                                        <FormControl style={{ width: 210, marginLeft: 20 }}>
                                            <InputLabel>จังหวัด</InputLabel>
                                            <Select
                                                id="blood_type"
                                                onChange={handleChange}
                                                variant='outlined'
                                            >
                                            </Select>
                                        </FormControl>
                                        <FormControl style={{ width: 210, marginLeft: 20 }}>
                                            <InputLabel>อําเภอ</InputLabel>
                                            <Select
                                                id="blood_type"
                                                onChange={handleChange}
                                                variant='outlined'
                                            >
                                            </Select>
                                        </FormControl>
                                        <FormControl style={{ width: 210, marginLeft: 20 }}>
                                            <InputLabel>ตําบล</InputLabel>
                                            <Select
                                                id="blood_type"
                                                onChange={handleChange}
                                                variant='outlined'
                                            >
                                            </Select>
                                        </FormControl>
                                    </div>
                                    <div style={{ marginTop: 20 }}>
                                        <TextField
                                            value={phoneNumber}
                                            label="เบอร์โทรศัพท์"
                                            variant="outlined"
                                            onChange={handleChange}
                                            style={{ width: 325 }}
                                            inputProps={{ maxLength: 10 }}
                                        />
                                    </div>
                                </div>
                            </CardContent>
                        </Card>
                        <Card className={classes.card} style={{ marginTop: 10 }}>
                            <CardHeader title='ข้อมูลซักประวัติ' />
                            <CardContent>
                                <div>
                                    <div>
                                        <TextField
                                            label="เลือกวันที่"
                                            name="added"
                                            type="date"
                                            variant='outlined'
                                            InputLabelProps={{
                                                shrink: true,
                                            }}
                                            style={{ width: 200 }}
                                        />
                                        <TextField
                                            label="เลือกเวลา"
                                            name="added"
                                            type="time"
                                            variant='outlined'
                                            InputLabelProps={{
                                                shrink: true,
                                            }}
                                            style={{ marginLeft: 20, width: 150 }}
                                        />
                                    </div>
                                    <div style={{ marginTop: 20 }}>
                                        <TextField
                                            label="น้ำหนัก"
                                            variant="outlined"
                                            style={{ width: 90 }}
                                        />
                                        <TextField
                                            label="ส่วนสูง"
                                            variant="outlined"
                                            style={{ width: 90, marginLeft: 20 }}
                                        />
                                        <TextField
                                            label="อุณหภูมิ"
                                            variant="outlined"
                                            style={{ width: 95, marginLeft: 20 }}
                                        />
                                        <TextField
                                            label="ความดันโลหิต"
                                            variant="outlined"
                                            style={{ width: 120, marginLeft: 20 }}
                                        />
                                    </div>
                                    <div style={{ marginTop: 20 }}>
                                        <TextField
                                            label="อาการเบื้องต้น"
                                            variant="outlined"
                                            style={{ width: 370 }}
                                            multiline
                                            rows={3}
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

export default PatientHistory;

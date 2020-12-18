import React, { FC, useEffect } from 'react';
import { makeStyles, ThemeProvider, Theme } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import ArrowBackIcon from '@material-ui/icons/ArrowBack';
import SearchIcon from '@material-ui/icons/Search';

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
    card: {
        width: 600,
    },
    margin: {
        margin: theme.spacing(3),
    }
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
                        <Card className={classes.card} style={{ marginTop: 10 }}>
                            <CardContent>
                                <div className={classes.margin}>
                                    <div>
                                        <TextField
                                            label="เลขบัตรประชาชน 13 หลัก"
                                            required
                                            variant='outlined'
                                            style={{ width: 440 }}
                                        />
                                        <Button
                                            variant='contained'
                                            style={{ marginLeft: 10, height: 55 }}
                                        >
                                            <SearchIcon fontSize='large' />
                                        </Button>
                                    </div>
                                    <div style={{ marginTop: 40 }}>
                                        <TextField
                                            disabled
                                            label="ชื่อผู้ป่วย"
                                            defaultValue="Mr. NAME SURNAME"
                                            variant='outlined'
                                            style={{ width: 520 }}
                                        />
                                    </div>
                                    <div style={{ marginTop: 40 }}>
                                        <TextField
                                            label="เลือกวันที่"
                                            name="added"
                                            type="date"
                                            variant='outlined'
                                            InputLabelProps={{
                                                shrink: true,
                                            }}
                                            style={{ width: 245 }}
                                        />
                                        <TextField
                                            label="เลือกเวลา"
                                            name="added"
                                            type="time"
                                            variant='outlined'
                                            InputLabelProps={{
                                                shrink: true,
                                            }}
                                            style={{ width: 245, marginLeft: 30 }}
                                        />
                                    </div>
                                    <div style={{ marginTop: 40 }}>
                                        <TextField
                                            label="น้ำหนัก"
                                            variant="outlined"
                                            style={{ width: 153 }}
                                        />
                                        <TextField
                                            label="ส่วนสูง"
                                            variant="outlined"
                                            style={{ width: 153, marginLeft: 30 }}
                                        />
                                        <TextField
                                            label="ความดันโลหิต"
                                            variant="outlined"
                                            style={{ width: 153, marginLeft: 30 }}
                                        />
                                    </div>
                                    <div style={{ marginTop: 40 }}>
                                        <TextField
                                            label="อาการเบื้องต้น"
                                            variant="outlined"
                                            style={{ width: 520 }}
                                            multiline
                                            rows={5}
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

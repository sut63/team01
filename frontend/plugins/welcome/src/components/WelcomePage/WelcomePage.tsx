import React, { useEffect, FC } from 'react';
import { Button,Typography, Grid, Link } from '@material-ui/core';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles,Theme ,createStyles} from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Avatar from '@material-ui/core/Avatar';
//import {  Cookies  } from 'react-cookie/cjs';//cookie
//import { useCookies } from 'react-cookie/cjs';//cookie


import { Image1Base64Function } from '../../image/Image1';
import { Image2Base64Function } from '../../image/Image2';
import { Image3Base64Function } from '../../image/Image3';
import { Image4Base64Function } from '../../image/Image4';
import { Image5Base64Function } from '../../image/Image5';
import { Image6Base64Function } from '../../image/image6_home';

const HeaderCustom = {
  minHeight: '50px',
};

const useStyles = makeStyles((theme: Theme) =>
createStyles({
  root: {
    maxWidth: 325,
  },
  formControl: {
    margin: theme.spacing(3),
    width: 320,
  },
  
}),
);

export type ProfileProps = {
  name: string; 
  id: string;
  system: string;
  imgsut: string;
  linkto: string;
};


export function CardTeam({ name, id, system ,imgsut,linkto}: ProfileProps) {
  
  const classes = useStyles();
  return (
    <Grid item xs={12} md={3}>
      <Card className={classes.root}>
        <CardActionArea>
        <Link
        href = {linkto}
        >
          <CardMedia
            component="img"
            height="140"
            image={imgsut}  
          />
           </Link>
          <CardContent>
            <Typography gutterBottom variant="subtitle1" component="h2">
              <br/>{system}
              <br/>{id} {name}
            </Typography>
          </CardContent>
        </CardActionArea>
      </Card>
    </Grid>
  );
}

const WelcomePage: FC<{}> = () => {
  /*
  const [cookies, setCookie, removeCookie] = useCookies(['cookiename']);

  const Logout = async () => {

    removeCookie('Name', { path: '/' })
    removeCookie('Password', { path: '/' })
    removeCookie('Log', { path: '/' })
    removeCookie('Status', { path: '/' })

  }
  */
  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ระบบห้องยา`}>
      {/*
      <Link href = "/">
      <Button
                onClick={() => {
                  Logout();
                }}
                variant="contained"
              >
                 Logout
             </Button>
             </Link>
             */}

      </Header>
      <Content>
        <ContentHeader title="สมาชิกในกลุ่ม"></ContentHeader>
        <Grid container>
        
          <CardTeam name={"นายภัทรพงศ์ เจริญผล"} id={"B6105358"} system={"ระบบบันทึกการชำระค่ายา"} imgsut = {Image1Base64Function} linkto = "/createBill"></CardTeam>
          <CardTeam name={"นางสาวชัญญา วัฒนาชีพ"} id={"B6113902"} system={"ระบบสั่งจ่ายยา"} imgsut = {Image2Base64Function}linkto = "/create_prescription"></CardTeam>
          <CardTeam name={"นายอนรรฆ ศิริมงคลขจร"} id={"B6114664"} system={"ระบบสั่งซื้อยา"} imgsut = {Image3Base64Function}linkto = "/create_Patientrights"></CardTeam>
          <CardTeam name={"นายธเนศ สุคันธตูล"} id={"B6115821"} system={"ระบบบันทึกประวัติการแพ้ยา"} imgsut = {Image4Base64Function}linkto = "/Treatment"></CardTeam>
          <CardTeam name={"นายกันตชาติ เหล่ากอ"} id={"B6117078"} system={"ระบบบันทึกยา"} imgsut = {Image5Base64Function}linkto = "/Historytaking"></CardTeam>
          <CardTeam name={"นายดุลยวัต อินหล้า"} id={"B6117566"} system={"ระบบบันทึกการจ่ายยา"} imgsut = {Image6Base64Function}linkto = "/Doctorinfo"></CardTeam>
         
        </Grid>
      </Content>
    </Page>
  );
};

export default WelcomePage;

import React, { useEffect } from 'react';
import HomeIcon from '@material-ui/icons/Home';

import EventNoteOutlinedIcon from '@material-ui/icons/EventNoteOutlined';
import QueuePlayNextOutlinedIcon from '@material-ui/icons/QueuePlayNextOutlined';
import PostAddOutlinedIcon from '@material-ui/icons/PostAddOutlined';
import HomeWorkOutlinedIcon from '@material-ui/icons/HomeWorkOutlined';
import MeetingRoomOutlinedIcon from '@material-ui/icons/MeetingRoomOutlined';
import AddShoppingCartOutlinedIcon from '@material-ui/icons/AddShoppingCartOutlined';
import ReceiptOutlinedIcon from '@material-ui/icons/ReceiptOutlined';
import PersonOutlineOutlinedIcon from '@material-ui/icons/PersonOutlineOutlined';
import SearchRoundedIcon from '@material-ui/icons/SearchRounded';
import VpnKeyRoundedIcon from '@material-ui/icons/VpnKeyRounded';
import { Cookies } from 'react-cookie/cjs'; //cookie

import {
  Sidebar,
  SidebarItem,
  SidebarDivider,
  SidebarSpace,
  SidebarThemeToggle,
  SidebarPinButton,
} from '@backstage/core';

const cookies = new Cookies();

const resetUserData = async () => {
  cookies.set('ID', JSON.stringify(0), { path: '/' });
  cookies.set('Name', JSON.stringify(null), { path: '/' });
  cookies.set('PositionData', JSON.stringify(''), { path: '/' });
  cookies.set('StatusLogin', JSON.stringify('No'), { path: '/' });
};

if (cookies.get('StatusLogin') == undefined) {
  resetUserData();
} else {
  var sPositionData = cookies.get('PositionData');
}

export const AppSidebar = () => (
  <Sidebar>
    <SidebarDivider />
    {/* Global nav, not org-specific */}
    {cookies.get('StatusLogin') == 'Yes' ? (
      <SidebarItem
        icon={PersonOutlineOutlinedIcon}
        text={cookies.get('Name')}
      />
    ) : (
      <>
        <SidebarItem icon={HomeIcon} to="" text="หน้าหลัก" />
        <SidebarItem icon={VpnKeyRoundedIcon} to="Signin" text="เข้าสู่ระบบ" />
      </>
    )}
    {/* <SidebarItem icon={CreateComponentIcon} to="create" text="Create..." />
    <SidebarItem icon={CreateComponentIcon} to="welcome" text="Welcome" /> */}

    <SidebarDivider />

    {cookies.get('StatusLogin') == 'Yes' ? (
      sPositionData == 'DrugAllergy' ? (
        <>
          <SidebarItem
            icon={EventNoteOutlinedIcon}
            to="DrugAllergy"
            text="ประวัติการเเพ้ยา"
          />
          <SidebarItem
            icon={SearchRoundedIcon}
            to="/SearchDrugAllergy"
            text="ค้นหาประวัติการเเพ้"
          />
        </>
      ) : sPositionData == 'Order' ? (
        <>
          <SidebarItem
            icon={AddShoppingCartOutlinedIcon}
            to="Order"
            text="สั่งซื้อยา"
          />
          <SidebarItem
            icon={SearchRoundedIcon}
            to="/SearchOrder"
            text="ค้นหาการสั่งซื้อยาย้อนหลัง"
          />
        </>
      ) : sPositionData == 'Medicine' ? (
        <>
          <SidebarItem
            icon={HomeWorkOutlinedIcon}
            to="Medicine"
            text="บันทึกข้อมูลยา"
          />
          <SidebarItem
            icon={SearchRoundedIcon}
            to="/SearchMedicine"
            text="ค้นหาข้อมูลยา"
          />
        </>
      ) : sPositionData == 'Prescription' ? (
        <>
          <SidebarItem
            icon={QueuePlayNextOutlinedIcon}
            to="Prescription"
            text="สั่งจ่ายยา"
          />
          <SidebarItem
            icon={SearchRoundedIcon}
            to="/TablePrescription"
            text="ค้นหารายการการสั่งจ่ายยา"
          />
        </>
      ) : sPositionData == 'DispenseMedicine' ? (
        <>
          <SidebarItem
            icon={PostAddOutlinedIcon}
            to="DispenseMedicine"
            text="บันทึกการจ่ายยา"
          />
          <SidebarItem
            icon={SearchRoundedIcon}
            to="SearchDispenseMedicine"
            text="ค้นหาประวัติการจ่ายยา"
          />
        </>
      ) : sPositionData == 'Bill' ? (
        <>
          <SidebarItem icon={ReceiptOutlinedIcon} to="Bill" text="ชำระค่ายา" />
          <SidebarItem icon={SearchRoundedIcon} to="" text="ค้นหา------" />
        </>
      ) : null
    ) : null}

    {/* End global nav */}
    <SidebarSpace />
    {cookies.get('StatusLogin') == 'Yes' ? (
      <SidebarItem
        icon={MeetingRoomOutlinedIcon}
        text="ออกจากระบบ"
        onClick={() => {
          resetUserData();
          history.pushState('', '', '/');
          window.location.reload(false);
        }}
      />
    ) : null}
    <SidebarDivider />
    <SidebarThemeToggle />
    {/* <SidebarUserSettings  /> */}
    <SidebarPinButton />
  </Sidebar>
);

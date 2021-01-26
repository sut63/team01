import React from 'react';
import HomeIcon from '@material-ui/icons/Home';

import EventNoteOutlinedIcon from '@material-ui/icons/EventNoteOutlined';
import QueuePlayNextOutlinedIcon from '@material-ui/icons/QueuePlayNextOutlined';
import PostAddOutlinedIcon from '@material-ui/icons/PostAddOutlined';
import HomeWorkOutlinedIcon from '@material-ui/icons/HomeWorkOutlined';
import MeetingRoomOutlinedIcon from '@material-ui/icons/MeetingRoomOutlined';
import AddShoppingCartOutlinedIcon from '@material-ui/icons/AddShoppingCartOutlined';
import ReceiptOutlinedIcon from '@material-ui/icons/ReceiptOutlined';
import PersonOutlineOutlinedIcon from '@material-ui/icons/PersonOutlineOutlined';
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

if (cookies.get('StatusLogin') == undefined) {
  cookies.set('ID', JSON.stringify(0), { path: '/' });
  cookies.set('Name', JSON.stringify(null), { path: '/' });
  cookies.set('PositionData', JSON.stringify('/'), { path: '/' });
  cookies.set('StatusLogin', JSON.stringify('No'), { path: '/' });
} else {
  var sPositionData = cookies.get('PositionData');
}

export const AppSidebar = () => (
  <Sidebar>
    <SidebarDivider />
    {/* Global nav, not org-specific */}
    {cookies.get('StatusLogin') == 'No' ? (
      <SidebarItem icon={HomeIcon} to="" text="Home" />
    ) : (
      <SidebarItem
        icon={PersonOutlineOutlinedIcon}
        text={cookies.get('Name')}
      />
    )}
    {/* <SidebarItem icon={CreateComponentIcon} to="create" text="Create..." />
    <SidebarItem icon={CreateComponentIcon} to="welcome" text="Welcome" /> */}

    <SidebarDivider />

    {cookies.get('StatusLogin') == 'Yes' ? (
      sPositionData == 'DrugAllergy' ? (
        <SidebarItem
          icon={EventNoteOutlinedIcon}
          to="DrugAllergy"
          text="ประวัติการเเพ้ยา"
        />
      ) : sPositionData == 'Order' ? (
        <SidebarItem
          icon={AddShoppingCartOutlinedIcon}
          to="Order"
          text="สั่งซื้อยา"
        />
      ) : sPositionData == 'Medicine' ? (
        <SidebarItem
          icon={HomeWorkOutlinedIcon}
          to="Medicine"
          text="บันทึกข้อมูลยา"
        />
      ) : sPositionData == 'Prescription' ? (
        <SidebarItem
          icon={QueuePlayNextOutlinedIcon}
          to="Prescription"
          text="สั่งจ่ายยา"
        />
        
      ) : sPositionData == 'DispenseMedicine' ? (
        <SidebarItem
          icon={PostAddOutlinedIcon}
          to="DispenseMedicine"
          text="บันทึกการจ่ายยา"
        />
      ) : sPositionData == 'CreateBill' ? (
        <SidebarItem
          icon={ReceiptOutlinedIcon}
          to="CreateBill"
          text="ชำระค่ายา"
        />
      ) : null
    ) : null}

    {/* End global nav */}
    <SidebarSpace />
    {cookies.get('StatusLogin') == 'Yes' ? (
      <SidebarItem
        icon={MeetingRoomOutlinedIcon}
        text="ออกจากระบบ"
        onClick={() => {
          cookies.set('ID', JSON.stringify(0), { path: '/' });
          cookies.set('Name', JSON.stringify(null), { path: '/' });
          cookies.set('PositionData', JSON.stringify('/'), { path: '/' });
          cookies.set('StatusLogin', JSON.stringify('No'), { path: '/' });
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

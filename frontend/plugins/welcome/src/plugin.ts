import { createPlugin } from '@backstage/core';
import { Cookies } from 'react-cookie/cjs'; //cookie
import WelcomePage from './components/WelcomePage';
import SignIn from './components/SignIn';

import DrugAllergy from './components/DrugAllergy';
import create_prescription from './components/create_prescription';
import Medicine from './components/Medicine';
import DispenseMedicine from './components/DispenseMedicine';
import Order from './components/Order';
import CreateBill from './components/createBill';

const cookies = new Cookies();
const sPositionData = cookies.get('PositionData');
const sLogin = cookies.get('StatusLogin');

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    if (sLogin == 'No') {
      router.registerRoute('/', WelcomePage);
      router.registerRoute('/signin', SignIn);
    } else {
      if (sPositionData == 'DrugAllergy') {
        router.registerRoute('/DrugAllergy', DrugAllergy);
      }
      if (sPositionData == 'Order') {
        router.registerRoute('/Order', Order);
      }
      if (sPositionData == 'Medicine') {
        router.registerRoute('/Medicine', Medicine);
      }
      if (sPositionData == 'Prescription') {
        router.registerRoute('/Prescription', create_prescription);
      }
      if (sPositionData == 'DispenseMedicine') {
        router.registerRoute('/DispenseMedicine', DispenseMedicine);
      }
      if (sPositionData == 'CreateBill') {
        router.registerRoute('/CreateBill', CreateBill);
      }
    }
  },
});

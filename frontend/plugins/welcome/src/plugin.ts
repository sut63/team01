import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import SignIn from './components/SignIn';

import DrugAllergy from './components/DrugAllergy';
import SearchDrugAllergy from './components/SearchDrugAllergy/SearchDrugAllergy';
import create_prescription from './components/create_prescription';
import Medicine from './components/Medicine';
import DispenseMedicine from './components/DispenseMedicine';
import SearchDispenseMedicine from './components/SearchDispenseMedicine';
import Order from './components/Order';
import CreateBill from './components/createBill';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/Signin', SignIn);
    //
    router.registerRoute('/DrugAllergy', DrugAllergy);
    router.registerRoute('/SearchDrugAllergy', SearchDrugAllergy);
    //
    router.registerRoute('/Order', Order);
    //
    router.registerRoute('/Medicine', Medicine);
    //
    router.registerRoute('/Prescription', create_prescription);
    //
    router.registerRoute('/DispenseMedicine', DispenseMedicine);
    router.registerRoute('/SearchDispenseMedicine', SearchDispenseMedicine);
    //
    router.registerRoute('/Bill', CreateBill);
  },
});

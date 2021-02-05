import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import SignIn from './components/SignIn';

import DrugAllergy from './components/DrugAllergy';
import SearchDrugAllergy from './components/SearchDrugAllergy/SearchDrugAllergy';
import create_prescription from './components/create_prescription';
import Medicine from './components/Medicine';
import DispenseMedicine from './components/DispenseMedicine';
import SearchDispenseMedicine from './components/SearchDispenseMedicine';
import SearchMedicine from './components/SearchMedicine';
import Order from './components/Order';
import CreateBill from './components/createBill';
import table_presrciption from './components/table_prescription';
import SearchOrder from './components/Searchorder'
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
    router.registerRoute('/SearchOrder', SearchOrder);
    //
    router.registerRoute('/Medicine', Medicine);
    router.registerRoute('/SearchMedicine', SearchMedicine);
    //
    router.registerRoute('/Prescription', create_prescription);
    router.registerRoute('/TablePrescription', table_presrciption);
    //
    router.registerRoute('/DispenseMedicine', DispenseMedicine);
    router.registerRoute('/SearchDispenseMedicine', SearchDispenseMedicine);
    //
    router.registerRoute('/Bill', CreateBill);
  },
});

import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import WatchVideo from './components/WatchVideo'
import SignIn from './components/SignIn'
import DrugAllergy from './components/DrugAllergy'
import create_prescription from './components/create_prescription'
import Medicine from './components/Medicine'
import DispenseMedicine from './components/DispenseMedicine'
import Order from './components/Order'

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/watch_video', WatchVideo);
    router.registerRoute('/signin', SignIn);
    router.registerRoute('/drug_allergy', DrugAllergy);
    router.registerRoute('/create_prescription', create_prescription);
    router.registerRoute('/medicine', Medicine);
    router.registerRoute('/dispense_medicine', DispenseMedicine);
    router.registerRoute('/order', Order);
  },
});

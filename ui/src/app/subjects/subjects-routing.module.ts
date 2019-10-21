import { NgModule } from '@angular/core';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';

import { Routes, RouterModule } from '@angular/router';
import { ListComponent } from './list/list.component';
import { ViewComponent } from './view/view.component';



const routes: Routes = [{
  path: '',
  component: ListComponent
},
{
  path: ':id',
  component: ViewComponent
}];

@NgModule({
  imports: [RouterModule.forChild(routes), NgbModule],
  exports: [RouterModule]
})
export class SubjectsRoutingModule { }

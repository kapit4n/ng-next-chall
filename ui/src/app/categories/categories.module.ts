import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';


import { CategoriesRoutingModule } from './categories-routing.module';
import { ListComponent } from './list/list.component';
import { CategoriesService } from './categories.service';
import { ComponentsModule } from '../components/components.module'

@NgModule({
  declarations: [
    ListComponent
  ],
  imports: [
    CommonModule,
    CategoriesRoutingModule,
    ComponentsModule
  ],
  providers: [CategoriesService]
})
export class CategoriesModule { }

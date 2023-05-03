import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { CategoriesRoutingModule } from './categories-routing.module';
import { ListComponent } from './list/list.component';
import { MatCardModule } from '@angular/material/card'
import { CategoriesService } from './categories.service';


@NgModule({
  declarations: [
    ListComponent
  ],
  imports: [
    CommonModule,
    CategoriesRoutingModule,
    MatCardModule,
  ],
  providers: [CategoriesService]
})
export class CategoriesModule { }

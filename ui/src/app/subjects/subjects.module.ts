import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';

import { SubjectsRoutingModule } from './subjects-routing.module';
import { ListComponent } from './list/list.component';
import { ViewComponent } from './view/view.component';

import { SubjectsService } from './subjects.service'

@NgModule({
  declarations: [ListComponent, ViewComponent],
  imports: [
    CommonModule,
    SubjectsRoutingModule,
    HttpClientModule
  ],
  providers: [SubjectsService]
})
export class SubjectsModule { }

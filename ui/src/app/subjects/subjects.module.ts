import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { SubjectsRoutingModule } from './subjects-routing.module';
import { ListComponent } from './list/list.component';
import { ViewComponent } from './view/view.component';

import { SubjectsService } from './subjects.service'

@NgModule({
  declarations: [ListComponent, ViewComponent],
  imports: [
    CommonModule,
    SubjectsRoutingModule
  ],
  providers: [SubjectsService]
})
export class SubjectsModule { }

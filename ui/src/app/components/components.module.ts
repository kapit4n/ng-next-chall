import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatCardModule } from '@angular/material/card'
import { CardCustomComponent } from './card-custom/card-custom.component'

import { CardMtComponent } from './card-mt/card-mt.component'

@NgModule({
  declarations: [
    CardCustomComponent,
    CardMtComponent
  ],
  imports: [
    CommonModule,
    MatCardModule
  ],
  exports: [
    CardMtComponent,
    CardCustomComponent
  ]
})
export class ComponentsModule { }

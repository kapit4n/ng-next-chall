import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatCardModule } from '@angular/material/card'
import { CardCustomComponent } from './card-custom/card-custom.component'

import { CardMtComponent } from './card-mt/card-mt.component';
import { CardComponent } from './card/card.component'

@NgModule({
  declarations: [
    CardCustomComponent,
    CardMtComponent,
    CardComponent
  ],
  imports: [
    CommonModule,
    MatCardModule
  ],
  exports: [
    CardComponent
  ]
})
export class ComponentsModule { }

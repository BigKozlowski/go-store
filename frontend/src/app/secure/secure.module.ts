import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SecureComponent } from './secure.component';
import { MenuComponent } from './menu/menu.component';
import { NavComponent } from './nav/nav.component';



@NgModule({
  declarations: [
    SecureComponent,
    MenuComponent,
    NavComponent
  ],
  imports: [
    CommonModule
  ]
})
export class SecureModule { }

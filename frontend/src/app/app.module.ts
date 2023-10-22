import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { PublicModule } from './public/public.module';
import { SecureModule } from './secure/secure.module';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    PublicModule,
    SecureModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

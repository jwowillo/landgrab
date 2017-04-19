import 'package:angular2/platform/browser.dart';

import 'package:landgrab/service/players.dart';
import 'package:landgrab/service/rules.dart';
import 'package:landgrab/service/schema.dart';

import 'package:landgrab/component/app/component.dart';

void main() {
  bootstrap(AppComponent, [PlayersService, RulesService, SchemaService]);
}

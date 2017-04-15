import 'package:angular2/core.dart';

import 'package:landgrab/model/rules.dart';

@Component(
  selector: 'rules',
  templateUrl: 'template.html',
)
class RulesComponent {
  @Input()
  Rules rules;
}

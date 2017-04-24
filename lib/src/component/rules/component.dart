import 'package:angular2/core.dart';

import 'package:landgrab/landgrab.dart';

@Component(
  selector: 'rules',
  templateUrl: 'template.html',
)
class RulesComponent {
  @Input()
  Rules rules;
}

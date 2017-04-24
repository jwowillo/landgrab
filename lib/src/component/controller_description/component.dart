import 'package:angular2/core.dart';

import 'package:landgrab/landgrab.dart';

@Component(
  selector: 'controller-description',
  templateUrl: 'template.html',
  directives: const [MethodDescriptionComponent],
)
class ControllerDescriptionComponent {
  @Input()
  ControllerDescription description;
}

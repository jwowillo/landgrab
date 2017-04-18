import 'package:angular2/core.dart';

import 'package:landgrab/component/method_description/component.dart';
import 'package:landgrab/model/api.dart';

@Component(
  selector: 'controller-description',
  templateUrl: 'template.html',
  directives: const [MethodDescriptionComponent],
)
class ControllerDescriptionComponent {
  @Input()
  ControllerDescription description;
}
